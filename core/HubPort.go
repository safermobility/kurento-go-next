// Code generated by kurento-go-generator. DO NOT EDIT.

package core

type IHubPort interface {
}

// This `MediaElement` specifies a connection with a `Hub`
type HubPort struct {
	MediaElement
}

type HubPort_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Hub Hub `json:"hub"`
}

func (HubPort_builder) GetTypeName() string {
	return "HubPort"
}
