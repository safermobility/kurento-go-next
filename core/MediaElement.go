// Code generated by kurento-go-generator. DO NOT EDIT.

package core

import (
	"context"
	"fmt"

	"github.com/safermobility/kurento-go-next/v6"
)

type IMediaElement interface {
	GetSourceConnections(context.Context, *MediaElementGetSourceConnectionsParams) ([]ElementConnectionData, error)
	GetSinkConnections(context.Context, *MediaElementGetSinkConnectionsParams) ([]ElementConnectionData, error)
	Connect(context.Context, *MediaElementConnectParams) error
	Disconnect(context.Context, *MediaElementDisconnectParams) error
	SetAudioFormat(context.Context, *MediaElementSetAudioFormatParams) error
	SetVideoFormat(context.Context, *MediaElementSetVideoFormatParams) error
	GetGstreamerDot(context.Context, *MediaElementGetGstreamerDotParams) (string, error)
	SetOutputBitrate(context.Context, *MediaElementSetOutputBitrateParams) error
	GetStats(context.Context, *MediaElementGetStatsParams) (map[string]Stats, error)
	IsMediaFlowingIn(context.Context, *MediaElementIsMediaFlowingInParams) (bool, error)
	IsMediaFlowingOut(context.Context, *MediaElementIsMediaFlowingOutParams) (bool, error)
	IsMediaTranscoding(context.Context, *MediaElementIsMediaTranscodingParams) (bool, error)
}

// The basic building block of the media server, that can be interconnected inside a pipeline.
// <p>
// A `MediaElement` is a module that encapsulates a specific media
// capability, and that is able to exchange media with other MediaElements
// through an internal element called <b>pad</b>.
// </p>
// <p>
// A pad can be defined as an input or output interface. Input pads are called
// sinks, and it's where the media elements receive media from other media
// elements. Output interfaces are called sources, and it's the pad used by the
// media element to feed media to other media elements. There can be only one
// sink pad per media element. On the other hand, the number of source pads is
// unconstrained. This means that a certain media element can receive media only
// from one element at a time, while it can send media to many others. Pads are
// created on demand, when the connect method is invoked. When two media elements
// are connected, one media pad is created for each type of media connected. For
// example, if you connect AUDIO and VIDEO between two media elements, each one
// will need to create two new pads: one for AUDIO and one for VIDEO.
// </p>
// <p>
// When media elements are connected, it can be the case that the encoding
// required in both input and output pads is not the same, and thus it needs to
// be transcoded. This is something that is handled transparently by the
// MediaElement internals, but such transcoding has a toll in the form of a
// higher CPU load, so connecting MediaElements that need media encoded in
// different formats is something to consider as a high load operation. The event
// `MediaTranscodingStateChanged` allows to inform the client application of
// whether media transcoding is being enabled or not inside any MediaElement
// object.
// </p>
type MediaElement struct {
	MediaObject

	// Minimum video bandwidth for transcoding.
	// @deprecated Deprecated due to a typo. Use :rom:meth:`minOutputBitrate` instead of this function.
	MinOuputBitrate int

	// Minimum video bitrate for transcoding.
	// <ul>
	// <li>Unit: bps (bits per second).</li>
	// <li>Default: 0.</li>
	// </ul>
	//
	MinOutputBitrate int

	// Maximum video bandwidth for transcoding.
	// @deprecated Deprecated due to a typo. Use :rom:meth:`maxOutputBitrate` instead of this function.
	MaxOuputBitrate int

	// Maximum video bitrate for transcoding.
	// <ul>
	// <li>Unit: bps (bits per second).</li>
	// <li>Default: MAXINT.</li>
	// <li>0 = unlimited.</li>
	// </ul>
	//
	MaxOutputBitrate int
}

type MediaElement_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (MediaElement_builder) GetTypeName() string {
	return "MediaElement"
}

type MediaElementGetSourceConnectionsParams struct {
	MediaType   MediaType `json:"MediaType"`
	Description string    `json:"Description"`
}

func (MediaElementGetSourceConnectionsParams) OperationName() string {
	return "getSourceConnections"
}

// Gets information about the sink pads of this media element.
// <p>
// Since sink pads are the interface through which a media element gets it's
// media, whatever is connected to an element's sink pad is formally a source of
// media. Media can be filtered by type, or by the description given to the pad
// though which both elements are connected.
// </p>
//
// Returns:
// // A list of the connections information that are sending media to this element. The list will be empty if no sources are found.
func (elem *MediaElement) GetSourceConnections(ctx context.Context, params *MediaElementGetSourceConnectionsParams) ([]ElementConnectionData, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // A list of the connections information that are sending media to this element. The list will be empty if no sources are found.
	value, err := kurento.CallSimple[[]ElementConnectionData](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}

type MediaElementGetSinkConnectionsParams struct {
	MediaType   MediaType `json:"MediaType"`
	Description string    `json:"Description"`
}

func (MediaElementGetSinkConnectionsParams) OperationName() string {
	return "getSinkConnections"
}

// Gets information about the source pads of this media element.
// <p>
// Since source pads connect to other media element's sinks, this is formally the
// sink of media from the element's perspective. Media can be filtered by type,
// or by the description given to the pad though which both elements are
// connected.
// </p>
//
// Returns:
// // A list of the connections information that are receiving media from this element. The list will be empty if no sources are found.
func (elem *MediaElement) GetSinkConnections(ctx context.Context, params *MediaElementGetSinkConnectionsParams) ([]ElementConnectionData, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // A list of the connections information that are receiving media from this element. The list will be empty if no sources are found.
	value, err := kurento.CallSimple[[]ElementConnectionData](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}

type MediaElementConnectParams struct {
	Sink                   IMediaElement `json:"Sink"`
	MediaType              MediaType     `json:"MediaType"`
	SourceMediaDescription string        `json:"SourceMediaDescription"`
	SinkMediaDescription   string        `json:"SinkMediaDescription"`
}

func (MediaElementConnectParams) OperationName() string {
	return "connect"
}

// Connects two elements, with the media flowing from left to right.
// <p>
// The element that invokes the connect will be the source of media, creating one
// sink pad for each type of media connected. The element given as parameter to
// the method will be the sink, and it will create one sink pad per media type
// connected.
// </p>
// <p>
// If otherwise not specified, all types of media are connected by default
// (AUDIO, VIDEO and DATA). It is recommended to connect the specific types of
// media if not all of them will be used. For this purpose, the connect method
// can be invoked more than once on the same two elements, but with different
// media types.
// </p>
// <p>
// The connection is unidirectional. If a bidirectional connection is desired,
// the position of the media elements must be inverted. For instance,
// webrtc1.connect(webrtc2) is connecting webrtc1 as source of webrtc2. In order
// to create a WebRTC one-2one conversation, the user would need to specify the
// connection on the other direction with webrtc2.connect(webrtc1).
// </p>
// <p>
// Even though one media element can have one sink pad per type of media, only
// one media element can be connected to another at a given time. If a media
// element is connected to another, the former will become the source of the sink
// media element, regardless whether there was another element connected or not.
// </p>
func (elem *MediaElement) Connect(ctx context.Context, params *MediaElementConnectParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type MediaElementDisconnectParams struct {
	Sink                   IMediaElement `json:"Sink"`
	MediaType              MediaType     `json:"MediaType"`
	SourceMediaDescription string        `json:"SourceMediaDescription"`
	SinkMediaDescription   string        `json:"SinkMediaDescription"`
}

func (MediaElementDisconnectParams) OperationName() string {
	return "disconnect"
}

// Disconnects two media elements. This will release the source pads of the source media element, and the sink pads of the sink media element.
func (elem *MediaElement) Disconnect(ctx context.Context, params *MediaElementDisconnectParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type MediaElementSetAudioFormatParams struct {
	Caps AudioCaps `json:"Caps"`
}

func (MediaElementSetAudioFormatParams) OperationName() string {
	return "setAudioFormat"
}

// Set the type of data for the audio stream.
// <p>
// MediaElements that do not support configuration of audio capabilities will
// throw a MEDIA_OBJECT_ILLEGAL_PARAM_ERROR exception.
// </p>
// <p>
// NOTE: This method is not implemented yet by the Media Server to do anything
// useful.
// </p>
func (elem *MediaElement) SetAudioFormat(ctx context.Context, params *MediaElementSetAudioFormatParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type MediaElementSetVideoFormatParams struct {
	Caps VideoCaps `json:"Caps"`
}

func (MediaElementSetVideoFormatParams) OperationName() string {
	return "setVideoFormat"
}

// Set the type of data for the video stream.
// <p>
// MediaElements that do not support configuration of video capabilities will
// throw a MEDIA_OBJECT_ILLEGAL_PARAM_ERROR exception
// </p>
// <p>
// NOTE: This method is not implemented yet by the Media Server to do anything
// useful.
// </p>
func (elem *MediaElement) SetVideoFormat(ctx context.Context, params *MediaElementSetVideoFormatParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type MediaElementGetGstreamerDotParams struct {
	Details GstreamerDotDetails `json:"Details"`
}

func (MediaElementGetGstreamerDotParams) OperationName() string {
	return "getGstreamerDot"
}

// Return a .dot file describing the topology of the media element.
// <p>The element can be queried for certain type of data:</p>
// <ul>
// <li>SHOW_ALL: default value</li>
// <li>SHOW_CAPS_DETAILS</li>
// <li>SHOW_FULL_PARAMS</li>
// <li>SHOW_MEDIA_TYPE</li>
// <li>SHOW_NON_DEFAULT_PARAMS</li>
// <li>SHOW_STATES</li>
// <li>SHOW_VERBOSE</li>
// </ul>
//
// Returns:
// // The dot graph.
func (elem *MediaElement) GetGstreamerDot(ctx context.Context, params *MediaElementGetGstreamerDotParams) (string, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // The dot graph.
	value, err := kurento.CallSimple[string](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}

type MediaElementSetOutputBitrateParams struct {
	Bitrate int `json:"Bitrate"`
}

func (MediaElementSetOutputBitrateParams) OperationName() string {
	return "setOutputBitrate"
}

// @deprecated
// Allows change the target bitrate for the media output, if the media is encoded using VP8 or H264. This method only works if it is called before the media starts to flow.
func (elem *MediaElement) SetOutputBitrate(ctx context.Context, params *MediaElementSetOutputBitrateParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type MediaElementGetStatsParams struct {
	MediaType MediaType `json:"MediaType"`
}

func (MediaElementGetStatsParams) OperationName() string {
	return "getStats"
}

// Gets the statistics related to an endpoint. If no media type is specified, it returns statistics for all available types.
// Returns:
// // Delivers a successful result in the form of a RTC stats report. A RTC stats report represents a map between strings, identifying the inspected objects (RTCStats.id), and their corresponding RTCStats objects.
func (elem *MediaElement) GetStats(ctx context.Context, params *MediaElementGetStatsParams) (map[string]Stats, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // Delivers a successful result in the form of a RTC stats report. A RTC stats report represents a map between strings, identifying the inspected objects (RTCStats.id), and their corresponding RTCStats objects.
	value, err := kurento.CallSimple[map[string]Stats](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}

type MediaElementIsMediaFlowingInParams struct {
	MediaType            MediaType `json:"MediaType"`
	SinkMediaDescription string    `json:"SinkMediaDescription"`
}

func (MediaElementIsMediaFlowingInParams) OperationName() string {
	return "isMediaFlowingIn"
}

// This method indicates whether the media element is receiving media of a certain type. The media sink pad can be identified individually, if needed. It is only supported for AUDIO and VIDEO types, raising a MEDIA_OBJECT_ILLEGAL_PARAM_ERROR otherwise. If the pad indicated does not exist, if will return false.
// Returns:
// // TRUE if there is media, FALSE in other case.
func (elem *MediaElement) IsMediaFlowingIn(ctx context.Context, params *MediaElementIsMediaFlowingInParams) (bool, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // TRUE if there is media, FALSE in other case.
	value, err := kurento.CallSimple[bool](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}

type MediaElementIsMediaFlowingOutParams struct {
	MediaType              MediaType `json:"MediaType"`
	SourceMediaDescription string    `json:"SourceMediaDescription"`
}

func (MediaElementIsMediaFlowingOutParams) OperationName() string {
	return "isMediaFlowingOut"
}

// This method indicates whether the media element is emitting media of a certain type. The media source pad can be identified individually, if needed. It is only supported for AUDIO and VIDEO types, raising a MEDIA_OBJECT_ILLEGAL_PARAM_ERROR otherwise. If the pad indicated does not exist, if will return false.
// Returns:
// // TRUE if there is media, FALSE in other case.
func (elem *MediaElement) IsMediaFlowingOut(ctx context.Context, params *MediaElementIsMediaFlowingOutParams) (bool, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // TRUE if there is media, FALSE in other case.
	value, err := kurento.CallSimple[bool](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}

type MediaElementIsMediaTranscodingParams struct {
	MediaType MediaType `json:"MediaType"`
	BinName   string    `json:"BinName"`
}

func (MediaElementIsMediaTranscodingParams) OperationName() string {
	return "isMediaTranscoding"
}

// Indicates whether this media element is actively transcoding between input and output pads. This operation is only supported for AUDIO and VIDEO media types, raising a MEDIA_OBJECT_ILLEGAL_PARAM_ERROR otherwise.
// The internal GStreamer processing bin can be indicated, if needed; if the bin doesn't exist, the return value will be FALSE.
// Returns:
// // TRUE if media is being transcoded, FALSE otherwise.
func (elem *MediaElement) IsMediaTranscoding(ctx context.Context, params *MediaElementIsMediaTranscodingParams) (bool, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // TRUE if media is being transcoded, FALSE otherwise.
	value, err := kurento.CallSimple[bool](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}
