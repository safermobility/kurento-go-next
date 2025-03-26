// Code generated by kurento-go-generator. DO NOT EDIT.

package elements

import (
	"context"
	"fmt"

	"github.com/safermobility/kurento-go-next"
	"github.com/safermobility/kurento-go-next/core"
)

type IMixer interface {
	Connect(media core.MediaType, source core.HubPort, sink core.HubPort) error
	Disconnect(media core.MediaType, source core.HubPort, sink core.HubPort) error
}

// A `Hub` that allows routing of video between arbitrary port pairs and mixing of audio among several ports
type Mixer struct {
	core.Hub
}

type Mixer_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	MediaPipeline core.MediaPipeline
}

func (Mixer_builder) GetTypeName() string {
	return "Mixer"
}

type MixerConnectParams struct {
	Media  core.MediaType `json:"Media"`
	Source core.HubPort   `json:"Source"`
	Sink   core.HubPort   `json:"Sink"`
}

func (MixerConnectParams) OperationName() string {
	return "connect"
}

// Connects each corresponding :rom:enum:`MediaType` of the given source port with the sink port.
func (elem *Mixer) Connect(ctx context.Context, params *MixerConnectParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type MixerDisconnectParams struct {
	Media  core.MediaType `json:"Media"`
	Source core.HubPort   `json:"Source"`
	Sink   core.HubPort   `json:"Sink"`
}

func (MixerDisconnectParams) OperationName() string {
	return "disconnect"
}

// Disonnects each corresponding :rom:enum:`MediaType` of the given source port from the sink port.
func (elem *Mixer) Disconnect(ctx context.Context, params *MixerDisconnectParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}
