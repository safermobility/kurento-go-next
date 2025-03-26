package kurento

//go:generate go run build/build.go
//go:generate go tool goimports -w core elements

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/wschannel"
)

type Client struct {
	c              *jrpc2.Client
	ws             *wschannel.Channel
	eventListeners *threadsafeSubscriberMap

	SessionID string
}

func New(url string, logger *log.Logger) (*Client, error) {
	ch, err := wschannel.Dial(url, &wschannel.DialOptions{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to start kurento websocket: %w", err)
	}

	subscribers := &threadsafeSubscriberMap{
		subscribers: make(map[string]map[string]map[string]EventHandler),
	}

	clientOpts := &jrpc2.ClientOptions{
		Logger: jrpc2.StdLogger(logger),
		OnNotify: func(r *jrpc2.Request) {
			if r.Method() == "onEvent" {
				var value struct {
					Value *Event
				}
				err := r.UnmarshalParams(&value)
				if err != nil {
					logger.Printf("error unmarshaling event '%s': '%s'", r.ParamString(), err)
					return
				}

				subscribers.handleEvent(value.Value)
			}
		},
	}
	c := jrpc2.NewClient(ch, clientOpts)

	// TODO: Implement ping - https://doc-kurento.readthedocs.io/en/latest/features/kurento_protocol.html#ping

	return &Client{
		c,
		ch,
		subscribers,
		"",
	}, nil
}

// func (c *Client) Call(ctx context.Context, method string, params map[string]interface{}) (*jrpc2.Response, error) {
// 	if c.SessionID != "" {
// 		params["sessionId"] = c.SessionID
// 	}
// 	r, err := c.c.Call(ctx, method, params)
//
// 	return r, err
// }

// SimpleResponse handles all Kurento responses except for "describe" type
type SimpleResponse[T any] struct {
	SessionID string
	Value     T
}

func CallSimple[T any](ctx context.Context, c *Client, params Request) (T, error) {
	params.SetSessionID(c.SessionID)
	var response SimpleResponse[T]
	err := c.c.CallResult(ctx, params.GetMethod(), params, &response)
	if err != nil {
		return response.Value, err
	}
	if sid := response.SessionID; sid != "" {
		c.SessionID = sid
	}
	return response.Value, nil
}

func (c *Client) Shutdown() {
	// TODO: What else has to happen here? Do we need to call both of these?
	c.c.Close()
	c.ws.Close()
}
