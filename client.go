package kurento

//go:generate go run build/build.go
//go:generate go tool goimports -w core elements

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"time"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/wschannel"
)

type Client struct {
	c              *jrpc2.Client
	ws             *wschannel.Channel
	logger         *slog.Logger
	eventListeners *threadsafeSubscriberMap

	SessionID string
}

func New(url string, logger *slog.Logger) (*Client, error) {
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
		Logger: func(text string) {
			var pcs [1]uintptr
			runtime.Callers(3, pcs[:]) // skip [Callers, this function, option wrapper]
			r := slog.NewRecord(time.Now(), slog.LevelInfo, text, pcs[0])
			_ = logger.Handler().Handle(context.Background(), r)
		},
		OnNotify: func(r *jrpc2.Request) {
			if r.Method() == "onEvent" {
				var value struct {
					Value *Event
				}
				err := r.UnmarshalParams(&value)
				if err != nil {
					logger.LogAttrs(
						context.Background(),
						slog.LevelError,
						"error unmarshaling event",
						slog.String("payload", r.ParamString()),
						slog.Any("error", err),
					)
					return
				}

				subscribers.handleEvent(value.Value)
				return
			}
			logger.LogAttrs(
				context.Background(),
				slog.LevelDebug,
				"unrecognized notification",
				slog.Any("payload", r),
			)
		},
	}
	c := jrpc2.NewClient(ch, clientOpts)

	// TODO: Implement ping - https://doc-kurento.readthedocs.io/en/latest/features/kurento_protocol.html#ping

	return &Client{
		c,
		ch,
		logger,
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
	c.c.Close()
	c.ws.Close()
}
