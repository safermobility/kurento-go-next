// Code generated by kurento-go-generator. DO NOT EDIT.

package elements

import "github.com/safermobility/kurento-go-next/v7/core"

type IComposite interface {
}

// A `Hub` that mixes the :rom:attr:`MediaType.AUDIO` stream of its connected sources and constructs a grid with the :rom:attr:`MediaType.VIDEO` streams of its connected sources into its sink
type Composite struct {
	core.Hub
}

type Composite_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	MediaPipeline core.MediaPipeline
}

func (Composite_builder) GetTypeName() string {
	return "Composite"
}
