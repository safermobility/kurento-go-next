// Code generated by kurento-go-generator. DO NOT EDIT.

package core

// Codec used for transmission of video.
type VideoCodec string

// Implement fmt.Stringer interface
func (t VideoCodec) String() string {
	return string(t)
}

const (
	VIDEOCODEC_VP8  VideoCodec = "VP8"
	VIDEOCODEC_H264 VideoCodec = "H264"
	VIDEOCODEC_RAW  VideoCodec = "RAW"
)
