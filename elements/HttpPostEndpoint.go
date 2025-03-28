// Code generated by kurento-go-generator. DO NOT EDIT.

package elements

import "github.com/safermobility/kurento-go-next/v6/core"

type IHttpPostEndpoint interface {
}

// An `HttpPostEndpoint` contains SINK pads for AUDIO and VIDEO, which provide access to an HTTP file upload function
//
// This type of endpoint provide unidirectional communications. Its `MediaSources <MediaSource>` are accessed through the HTTP POST method.
type HttpPostEndpoint struct {
	HttpEndpoint
}

type HttpPostEndpoint_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	MediaPipeline        core.MediaPipeline
	DisconnectionTimeout int
	UseEncodedMedia      bool
}

func (HttpPostEndpoint_builder) GetTypeName() string {
	return "HttpPostEndpoint"
}
