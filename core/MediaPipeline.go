// Code generated by kurento-go-generator. DO NOT EDIT.

package core

import (
	"context"
	"fmt"

	"github.com/safermobility/kurento-go-next"
)

type IMediaPipeline interface {
	GetGstreamerDot(details GstreamerDotDetails) (string, error)
}

// A pipeline is a container for a collection of `MediaElements<MediaElement>` and `MediaMixers<MediaMixer>`.
// It offers the methods needed to control the creation and connection of elements inside a certain pipeline.
type MediaPipeline struct {
	MediaObject

	// If statistics about pipeline latency are enabled for all mediaElements
	LatencyStats bool
}

type MediaPipeline_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (MediaPipeline_builder) GetTypeName() string {
	return "MediaPipeline"
}

type MediaPipelineGetGstreamerDotParams struct {
	Details GstreamerDotDetails `json:"Details"`
}

func (MediaPipelineGetGstreamerDotParams) OperationName() string {
	return "getGstreamerDot"
}

// Returns a string in dot (graphviz) format that represents the gstreamer elements inside the pipeline
// Returns:
// // The dot graph.
func (elem *MediaPipeline) GetGstreamerDot(ctx context.Context, params *MediaPipelineGetGstreamerDotParams) (string, error) {
	request := kurento.BuildInvoke(elem.Id, params)

	// // The dot graph.
	value, err := kurento.CallSimple[string](ctx, elem.GetConnection(), request)
	if err != nil {
		err = fmt.Errorf("rpc error: %w", err)
	}
	return value, err

}
