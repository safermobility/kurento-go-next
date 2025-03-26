package kurento

import (
	"context"
	"fmt"
	"sync"
)

type EventHandler func(map[string]interface{})

type threadsafeSubscriberMap struct {
	subscribers map[string]map[string]map[string]EventHandler // eventName -> objectId -> handlerId -> handler.
	lock        sync.RWMutex
}

type Event struct {
	Object string
	Type   string
	Data   map[string]any // TODO: Can we make types for all of these?
}

func (t *threadsafeSubscriberMap) handleEvent(e *Event) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if eventHandlers, ok := t.subscribers[e.Type]; ok {
		if objectHandlers, ok := eventHandlers[e.Object]; ok {
			for _, handler := range objectHandlers {
				handler(e.Data)
			}
		}
	}
}

type Subscription struct {
	// For subscribing
	EventType string `json:"type,omitempty"`
	// For unsubscribing
	SubscriptionID string `json:"subscription,omitempty"`
	// For both
	ObjectID  string `json:"object"`
	SessionID string `json:"sessionId"`
}

func (c *Client) Subscribe(ctx context.Context, event, objectId string, handler EventHandler) (string, error) {
	var oh map[string]map[string]EventHandler
	var ok bool

	c.eventListeners.lock.Lock()
	defer c.eventListeners.lock.Unlock()

	info := &Subscription{
		EventType: event,
		ObjectID:  objectId,
		SessionID: c.SessionID,
	}
	var resp SimpleResponse[string]
	err := c.c.CallResult(ctx, "subscribe", info, &resp)
	if err != nil {
		return resp.Value, fmt.Errorf("unable to subscribe to event '%s' on '%s': %w", info.EventType, info.ObjectID, err)
	}

	if oh, ok = c.eventListeners.subscribers[info.EventType]; !ok {
		c.eventListeners.subscribers[info.EventType] = make(map[string]map[string]EventHandler)
		oh = c.eventListeners.subscribers[info.EventType]
	}

	var he map[string]EventHandler
	if he, ok = oh[info.ObjectID]; !ok {
		oh[info.ObjectID] = make(map[string]EventHandler)
		he = oh[info.ObjectID]
	}

	he[resp.Value] = handler
	return resp.Value, nil
}

func (c *Client) Unsubscribe(ctx context.Context, event, objectId, handlerId string) error {
	c.eventListeners.lock.Lock()
	defer c.eventListeners.lock.Unlock()

	info := &Subscription{
		SubscriptionID: handlerId,
		ObjectID:       objectId,
		SessionID:      c.SessionID,
	}
	var resp SimpleResponse[any]
	err := c.c.CallResult(ctx, "unsubscribe", info, &resp)
	if err != nil {
		return fmt.Errorf("unable to unsubscribe from '%s' on '%s': %w", info.SubscriptionID, info.ObjectID, err)
	}

	if oh, ok := c.eventListeners.subscribers[event]; ok {
		delete(oh[objectId], handlerId)
	}

	return nil
}
