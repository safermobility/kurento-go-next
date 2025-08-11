package kurento

import (
	"context"
	"fmt"
	"sync"
)

type EventHandler func(map[string]interface{})

type threadsafeSubscriberMap struct {
	subscribers map[string]map[string]EventHandler // eventName+objectId -> handlerId -> handler.
	// Lock per type+object name
	locks   map[string]*sync.RWMutex
	mapLock sync.Mutex
}

type Event struct {
	Object string
	Type   string
	Data   map[string]any // TODO: Can we make types for all of these?
}

func (t *threadsafeSubscriberMap) getLockFor(key string) *sync.RWMutex {
	t.mapLock.Lock()
	defer t.mapLock.Unlock()

	if ret, found := t.locks[key]; found {
		return ret
	}

	ret := &sync.RWMutex{}
	t.locks[key] = ret
	return ret
}

func (t *threadsafeSubscriberMap) handleEvent(e *Event) {
	name := e.Type + ":" + e.Object
	lock := t.getLockFor(name)
	lock.RLock()
	defer lock.RUnlock()
	if eventHandlers, ok := t.subscribers[name]; ok {
		for _, handler := range eventHandlers {
			handler(e.Data)
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
	var ok bool

	name := event + ":" + objectId
	lock := c.eventListeners.getLockFor(name)
	lock.Lock()
	defer lock.Unlock()

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

	var he map[string]EventHandler
	if he, ok = c.eventListeners.subscribers[name]; !ok {
		c.eventListeners.subscribers[name] = make(map[string]EventHandler)
		he = c.eventListeners.subscribers[name]
	}

	he[resp.Value] = handler
	return resp.Value, nil
}

func (c *Client) Unsubscribe(ctx context.Context, event, objectId, handlerId string) error {
	name := event + ":" + objectId
	lock := c.eventListeners.getLockFor(name)
	lock.Lock()
	defer lock.Unlock()

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

	if oh, ok := c.eventListeners.subscribers[name]; ok {
		delete(oh, handlerId)
	}

	return nil
}
