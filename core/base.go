package core

import (
	"context"

	"github.com/safermobility/kurento-go-next/v6"
)

type IMediaObject interface {
	// Each media object should be able to create another object
	// Those options are sent to getConstructorParams
	Create(context.Context, IMediaObject, kurento.IMediaObjectBuilder) error

	Release(context.Context) error

	setConnection(*kurento.Client)

	// Set ID of the element
	setID(string)

	//Implement Stringer
	String() string

	setParent(IMediaObject)
	addChild(IMediaObject)
}

func (elem *MediaObject) Create(ctx context.Context, m IMediaObject, options kurento.IMediaObjectBuilder) error {
	request := kurento.BuildCreate(options)
	id, err := kurento.CallSimple[string](ctx, elem.connection, request)
	if err != nil {
		return err
	}

	m.setConnection(elem.connection)
	m.setID(id)
	m.setParent(elem)
	elem.addChild(m)

	return nil
}

func (elem *MediaObject) Release(ctx context.Context) error {
	request := kurento.BuildRelease(elem.String())
	_, err := kurento.CallSimple[any](ctx, elem.connection, request)
	if err != nil {
		return err
	}
	return nil
}

func (elem *MediaObject) Subscribe(ctx context.Context, event string, cb kurento.EventHandler) (string, error) {
	return elem.connection.Subscribe(ctx, event, elem.Id, cb)
}

func (elem *MediaObject) setConnection(c *kurento.Client) {
	elem.connection = c
}

func (elem *MediaObject) GetConnection() *kurento.Client {
	return elem.connection
}

func (elem *MediaObject) setParent(m IMediaObject) {
	elem.Parent = m
}

func (elem *MediaObject) addChild(m IMediaObject) {
	elem.Childs = append(elem.Childs, m)
}

func (m *MediaObject) setID(id string) {
	m.Id = id
}

func (m *MediaObject) String() string {
	return m.Id
}

// Create an object in memory that represents a remote object without creating it
func HydrateMediaObject(id string, parent IMediaObject, c *kurento.Client, elem IMediaObject) error {
	elem.setConnection(c)
	elem.setID(id)
	if parent != nil {
		parent.addChild(elem)
	}
	return nil
}
