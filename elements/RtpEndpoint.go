// Code generated by kurento-go-generator. DO NOT EDIT.

package elements

import "github.com/safermobility/kurento-go-next/v7/core"

type IRtpEndpoint interface {
}

// Endpoint that provides bidirectional content delivery capabilities through the
// RTP or SRTP protocols.
// <p>
// An `RtpEndpoint` contains paired sink and source `MediaPad`
// for audio and video. This endpoint inherits from `BaseRtpEndpoint`.
// </p>
// <p>
// In order to establish RTP/SRTP communications, peers first engage in an SDP
// Offer/Answer negotiation process, where one of the peers (the offerer) sends
// an SDP Offer, while the other peer (the answerer) responds with an SDP Answer.
// This endpoint can work in both roles.
// </p>
// <ul>
// <li>
// <b>As offerer</b>: The negotiation process is initiated by the media server.
// <ul>
// <li>
// Kurento generates the SDP Offer through the
// <code>generateOffer()</code> method. This offer must then be sent to the
// remote peer (the answerer) through the signaling channel.
// </li>
// <li>
// The remote peer process the SDP Offer, and generates an SDP Answer. This
// answer is then sent back to the media server.
// </li>
// <li>
// Upon receiving the SDP Answer, this endpoint must process it with the
// <code>processAnswer()</code> method.
// </li>
// </ul>
// </li>
// <li>
// <b>As answerer</b>: The negotiation process is initiated by the remote peer.
// <ul>
// <li>
// The remote peer, acting as offerer, generates an SDP Offer and sends it
// to this endpoint.
// </li>
// <li>
// This endpoint processes the SDP Offer with the
// <code>processOffer()</code> method. The result of this method will be a
// string, containing an SDP Answer.
// </li>
// <li>
// The SDP Answer must then be sent back to the offerer, so it can be
// processed by it.
// </li>
// </ul>
// </li>
// </ul>
// <p>
// In case of unidirectional connections (i.e. only one peer is going to send
// media), the process is simpler, as only the sender needs to process an SDP
// Offer. On top of the information about media codecs and types, the SDP must
// contain the IP of the remote peer, and the port where it will be listening.
// This way, the SDP can be mangled without needing to go through the exchange
// process, as the receiving peer does not need to process any answer.
// </p>
// <h2>Bitrate management</h2>
// <p>
// Check the documentation of `BaseRtpEndpoint` for detailed information
// about bitrate management.
// </p>
type RtpEndpoint struct {
	core.BaseRtpEndpoint
}

type RtpEndpoint_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	MediaPipeline core.MediaPipeline `json:"mediaPipeline"`
	Crypto        SDES               `json:"crypto,omitempty"`
	UseIpv6       bool               `json:"useIpv6,omitempty"`
}

func (RtpEndpoint_builder) GetTypeName() string {
	return "RtpEndpoint"
}
