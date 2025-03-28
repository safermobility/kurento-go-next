// Code generated by kurento-go-generator. DO NOT EDIT.

package core

type IEndpoint interface {
}

// Base interface for all end points.
// <p>
// An Endpoint is a `MediaElement` that allows Kurento to exchange
// media contents with external systems, supporting different transport protocols
// and mechanisms, such as RTP, WebRTC, HTTP(s), "file://" URLs, etc.
// </p>
// <p>
// An "Endpoint" may contain both sources and sinks for different media types,
// to provide bidirectional communication.
// </p>
type Endpoint struct {
	MediaElement
}

type Endpoint_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (Endpoint_builder) GetTypeName() string {
	return "Endpoint"
}
