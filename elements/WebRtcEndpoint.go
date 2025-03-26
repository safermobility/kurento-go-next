// Code generated by kurento-go-generator. DO NOT EDIT.

package elements

import (
	"context"
	"fmt"

	"github.com/safermobility/kurento-go-next/v6"
	"github.com/safermobility/kurento-go-next/v6/core"
)

type IWebRtcEndpoint interface {
	GatherCandidates() error
	AddIceCandidate(candidate IceCandidate) error
	CreateDataChannel(label string, ordered bool, maxPacketLifeTime int, maxRetransmits int, protocol string) error
	CloseDataChannel(channelId int) error
}

// Control interface for Kurento WebRTC endpoint.
// <p>
// This endpoint is one side of a peer-to-peer WebRTC communication, being the
// other peer a WebRTC capable browser -using the RTCPeerConnection API-, a
// native WebRTC app or even another Kurento Media Server.
// </p>
// <p>
// In order to establish a WebRTC communication, peers engage in an SDP
// negotiation process, where one of the peers (the offerer) sends an offer,
// while the other peer (the offeree) responds with an answer. This endpoint can
// function in both situations
// </p>
// <ul>
// <li>
// As offerer: The negotiation process is initiated by the media server
// <ul>
// <li>
// KMS generates the SDP offer through the
// <code>generateOffer</code> method. This <i>offer</i> must then be sent
// to the remote peer (the offeree) through the signaling channel, for
// processing.
// </li>
// <li>
// The remote peer processes the <i>offer</i>, and generates an
// <i>answer</i>. The <i>answer</i> is sent back to the media server.
// </li>
// <li>
// Upon receiving the <i>answer</i>, the endpoint must invoke the
// <code>processAnswer</code> method.
// </li>
// </ul>
// </li>
// <li>
// As offeree: The negotiation process is initiated by the remote peer
// <ul>
// <li>
// The remote peer, acting as offerer, generates an SDP <i>offer</i> and
// sends it to the WebRTC endpoint in Kurento.
// </li>
// <li>
// The endpoint will process the <i>offer</i> invoking the
// <code>processOffer</code> method. The result of this method will be a
// string, containing an SDP <i>answer</i>.
// </li>
// <li>
// The SDP <i>answer</i> must be sent back to the offerer, so it can be
// processed.
// </li>
// </ul>
// </li>
// </ul>
// <p>
// SDPs are sent without ICE candidates, following the Trickle ICE optimization.
// Once the SDP negotiation is completed, both peers proceed with the ICE
// discovery process, intended to set up a bidirectional media connection. During
// this process, each peer
// </p>
// <ul>
// <li>
// Discovers ICE candidates for itself, containing pairs of IPs and ports.
// </li>
// <li>
// ICE candidates are sent via the signaling channel as they are discovered, to
// the remote peer for probing.
// </li>
// <li>
// ICE connectivity checks are run as soon as the new candidate description,
// from the remote peer, is available.
// </li>
// </ul>
// <p>
// Once a suitable pair of candidates (one for each peer) is discovered, the
// media session can start. The harvesting process in Kurento, begins with the
// invocation of the <code>gatherCandidates</code> method. Since the whole
// Trickle ICE purpose is to speed-up connectivity, candidates are generated
// asynchronously. Therefore, in order to capture the candidates, the user must
// subscribe to the event <code>IceCandidateFound</code>. It is important that
// the event listener is bound before invoking <code>gatherCandidates</code>,
// otherwise a suitable candidate might be lost, and connection might not be
// established.
// </p>
// <p>
// It's important to keep in mind that WebRTC connection is an asynchronous
// process, when designing interactions between different MediaElements. For
// example, it would be pointless to start recording before media is flowing. In
// order to be notified of state changes, the application can subscribe to events
// generated by the WebRtcEndpoint. Following is a full list of events generated
// by WebRtcEndpoint:
// </p>
// <ul>
// <li>
// <code>IceComponentStateChange</code>: This event informs only about changes
// in the ICE connection state. Possible values are:
// <ul>
// <li><code>DISCONNECTED</code>: No activity scheduled</li>
// <li><code>GATHERING</code>: Gathering local candidates</li>
// <li><code>CONNECTING</code>: Establishing connectivity</li>
// <li><code>CONNECTED</code>: At least one working candidate pair</li>
// <li>
// <code>READY</code>: ICE concluded, candidate pair selection is now final
// </li>
// <li>
// <code>FAILED</code>: Connectivity checks have been completed, but media
// connection was not established
// </li>
// </ul>
// The transitions between states are covered in RFC5245. It could be said that
// it's network-only, as it only takes into account the state of the network
// connection, ignoring other higher level stuff, like DTLS handshake, RTCP
// flow, etc. This implies that, while the component state is
// <code>CONNECTED</code>, there might be no media flowing between the peers.
// This makes this event useful only to receive low-level information about the
// connection between peers. Even more, while other events might leave a
// graceful period of time before firing, this event fires immediately after
// the state change is detected.
// </li>
// <li>
// <code>IceCandidateFound</code>: Raised when a new candidate is discovered.
// ICE candidates must be sent to the remote peer of the connection. Failing to
// do so for some or all of the candidates might render the connection
// unusable.
// </li>
// <li>
// <code>IceGatheringDone</code>: Raised when the ICE gathering process is
// completed. This means that all candidates have already been discovered.
// </li>
// <li>
// <code>NewCandidatePairSelected</code>: Raised when a new ICE candidate pair
// gets selected. The pair contains both local and remote candidates being used
// for a component. This event can be raised during a media session, if a new
// pair of candidates with higher priority in the link are found.
// </li>
// <li><code>DataChannelOpen</code>: Raised when a data channel is open.</li>
// <li><code>DataChannelClose</code>: Raised when a data channel is closed.</li>
// </ul>
// <p>
// Registering to any of above events requires the application to provide a
// callback function. Each event provides different information, so it is
// recommended to consult the signature of the event listeners.
// </p>
// <p>
// Flow control and congestion management is one of the most important features
// of WebRTC. WebRTC connections start with the lowest bandwidth configured and
// slowly ramps up to the maximum available bandwidth, or to the higher limit of
// the exploration range in case no bandwidth limitation is detected. Notice that
// WebRtcEndpoints in Kurento are designed in a way that multiple WebRTC
// connections fed by the same stream share quality. When a new connection is
// added, as it requires to start with low bandwidth, it will cause the rest of
// connections to experience a transient period of degraded quality, until it
// stabilizes its bitrate. This doesn't apply when transcoding is involved.
// Transcoders will adjust their output bitrate based in bandwidth requirements,
// but it won't affect the original stream. If an incoming WebRTC stream needs to
// be transcoded, for whatever reason, all WebRtcEndpoints fed from transcoder
// output will share a separate quality than the ones connected directly to the
// original stream.
// </p>
// <p>
// Note that the default <strong>VideoSendBandwidth</strong> range of the
// endpoint is a VERY conservative one, and leads to a low maximum video quality.
// Most applications will probably want to increase this to higher values such as
// 2000 kbps (2 mbps).
// </p>
// <p>
// <strong>
// Check the extended documentation of these parameters in
// `SdpEndpoint`, `BaseRtpEndpoint`, and
// `RembParams`.
// </strong>
// </p>
// <ul>
// <li>
// Input bandwidth: Values used to inform remote peers about the bitrate that
// can be sent to this endpoint.
// <ul>
// <li>
// <strong>MinVideoRecvBandwidth</strong>: Minimum input bitrate, requested
// from WebRTC senders with REMB (Default: 30 Kbps).
// </li>
// <li>
// <strong>MaxAudioRecvBandwidth</strong> and
// <strong>MaxVideoRecvBandwidth</strong>: Maximum input bitrate, signaled
// in SDP Offers to WebRTC and RTP senders (Default: unlimited).
// </li>
// </ul>
// </li>
// <li>
// Output bandwidth: Values used to control bitrate of the video streams sent
// to remote peers. It is important to keep in mind that pushed bitrate depends
// on network and remote peer capabilities. Remote peers can also announce
// bandwidth limitation in their SDPs (through the
// <code>b={modifier}:{value}</code> attribute). Kurento will always enforce
// bitrate limitations specified by the remote peer over internal
// configurations.
// <ul>
// <li>
// <strong>MinVideoSendBandwidth</strong>: REMB override of minimum bitrate
// sent to WebRTC receivers (Default: 100 Kbps).
// </li>
// <li>
// <strong>MaxVideoSendBandwidth</strong>: REMB override of maximum bitrate
// sent to WebRTC receivers (Default: 500 Kbps).
// </li>
// <li>
// <strong>RembParams.rembOnConnect</strong>: Initial local REMB bandwidth
// estimation that gets propagated when a new endpoint is connected.
// </li>
// </ul>
// </li>
// </ul>
// <p>
// <strong>
// All bandwidth control parameters must be changed before the SDP negotiation
// takes place, and can't be changed afterwards.
// </strong>
// </p>
// <p>
// DataChannels allow other media elements that make use of the DataPad, to send
// arbitrary data. For instance, if there is a filter that publishes event
// information, it'll be sent to the remote peer through the channel. There is no
// API available for programmers to make use of this feature in the
// WebRtcElement. DataChannels can be configured to provide the following:
// </p>
// <ul>
// <li>Reliable or partially reliable delivery of sent messages</li>
// <li>In-order or out-of-order delivery of sent messages</li>
// </ul>
// <p>
// Unreliable, out-of-order delivery is equivalent to raw UDP semantics. The
// message may make it, or it may not, and order is not important. However, the
// channel can be configured to be <i>partially reliable</i> by specifying the
// maximum number of retransmissions or setting a time limit for retransmissions:
// the WebRTC stack will handle the acknowledgments and timeouts.
// </p>
// <p>
// The possibility to create DataChannels in a WebRtcEndpoint must be explicitly
// enabled when creating the endpoint, as this feature is disabled by default. If
// this is the case, they can be created invoking the createDataChannel method.
// The arguments for this method, all of them optional, provide the necessary
// configuration:
// </p>
// <ul>
// <li>
// <code>label</code>: assigns a label to the DataChannel. This can help
// identify each possible channel separately.
// </li>
// <li>
// <code>ordered</code>: specifies if the DataChannel guarantees order, which
// is the default mode. If maxPacketLifetime and maxRetransmits have not been
// set, this enables reliable mode.
// </li>
// <li>
// <code>maxPacketLifeTime</code>: The time window in milliseconds, during
// which transmissions and retransmissions may take place in unreliable mode.
// This forces unreliable mode, even if <code>ordered</code> has been
// activated.
// </li>
// <li>
// <code>maxRetransmits</code>: maximum number of retransmissions that are
// attempted in unreliable mode. This forces unreliable mode, even if
// <code>ordered</code> has been activated.
// </li>
// <li>
// <code>Protocol</code>: Name of the subprotocol used for data communication.
// </li>
// </ul>
type WebRtcEndpoint struct {
	core.BaseRtpEndpoint

	// Local network interfaces used for ICE gathering.
	// <p>
	// If you know which network interfaces should be used to perform ICE (for WebRTC
	// connectivity), you can define them here. Doing so has several advantages:
	// </p>
	// <ul>
	// <li>
	// The WebRTC ICE gathering process will be much quicker. Normally, it needs to
	// gather local candidates for all of the network interfaces, but this step can
	// be made faster if you limit it to only the interface that you know will
	// work.
	// </li>
	// <li>
	// It will ensure that the media server always decides to use the correct
	// network interface. With WebRTC ICE gathering it's possible that, under some
	// circumstances (in systems with virtual network interfaces such as
	// <code>docker0</code>) the ICE process ends up choosing the wrong local IP.
	// </li>
	// </ul>
	// <p>
	// <code>networkInterfaces</code> is a comma-separated list of network interface
	// names.
	// </p>
	// <p>Examples:</p>
	// <ul>
	// <li><code>networkInterfaces=eth0</code></li>
	// <li><code>networkInterfaces=eth0,enp0s25</code></li>
	// </ul>
	//
	NetworkInterfaces string

	// Enable ICE-TCP candidate gathering.
	// <p>
	// This setting enables or disables using TCP for ICE candidate gathering in the
	// underlying libnice library:
	// https://libnice.freedesktop.org/libnice/NiceAgent.html#NiceAgent--ice-tcp
	// </p>
	// <p>
	// You might want to disable ICE-TCP to potentially speed up ICE gathering by
	// avoiding TCP candidates in scenarios where they are not needed.
	// </p>
	// <p><code>iceTcp</code> is either 1 (ON) or 0 (OFF). Default: 1 (ON).</p>
	//
	IceTcp bool

	// STUN server IP address.
	// <p>The ICE process uses STUN to punch holes through NAT firewalls.</p>
	// <p>
	// <code>stunServerAddress</code> MUST be an IP address; domain names are NOT
	// supported.
	// </p>
	// <p>
	// You need to use a well-working STUN server. Use this to check if it works:<br />
	// https://webrtc.github.io/samples/src/content/peerconnection/trickle-ice/<br />
	// From that check, you should get at least one Server-Reflexive Candidate (type
	// <code>srflx</code>).
	// </p>
	//
	StunServerAddress string

	// Port of the STUN server
	StunServerPort int

	// TURN server URL.
	// <p>
	// When STUN is not enough to open connections through some NAT firewalls, using
	// TURN is the remaining alternative.
	// </p>
	// <p>
	// Note that TURN is a superset of STUN, so you don't need to configure STUN if
	// you are using TURN.
	// </p>
	// <p>The provided URL should follow one of these formats:</p>
	// <ul>
	// <li><code>user:password@ipaddress:port</code></li>
	// <li>
	// <code>user:password@ipaddress:port?transport=[udp|tcp|tls]</code>
	// </li>
	// </ul>
	// <p>
	// <code>ipaddress</code> MUST be an IP address; domain names are NOT supported.<br />
	// <code>transport</code> is OPTIONAL. Possible values: udp, tcp, tls. Default: udp.
	// </p>
	// <p>
	// You need to use a well-working TURN server. Use this to check if it works:<br />
	// https://webrtc.github.io/samples/src/content/peerconnection/trickle-ice/<br />
	// From that check, you should get at least one Server-Reflexive Candidate (type
	// <code>srflx</code>) AND one Relay Candidate (type <code>relay</code>).
	// </p>
	//
	TurnUrl string

	// External IPv4 address of the media server.
	// <p>
	// Forces all local IPv4 ICE candidates to have the given address. This is really
	// nothing more than a hack, but it's very effective to force a public IP address
	// when one is known in advance for the media server. In doing so, KMS will not
	// need a STUN or TURN server, but remote peers will still be able to contact it.
	// </p>
	// <p>
	// You can try using this setting if KMS is deployed on a publicly accessible
	// server, without NAT, and with a static public IP address. But if it doesn't
	// work for you, just go back to configuring a STUN or TURN server for ICE.
	// </p>
	// <p>
	// Only set this parameter if you know what you're doing, and you understand 100%
	// WHY you need it. For the majority of cases, you should just prefer to
	// configure a STUN or TURN server.
	// </p>
	// <p><code>externalIPv4</code> is a single IPv4 address.</p>
	// <p>Example:</p>
	// <ul>
	// <li><code>externalIPv4=198.51.100.1</code></li>
	// </ul>
	//
	ExternalIPv4 string

	// External IPv6 address of the media server.
	// <p>
	// Forces all local IPv6 ICE candidates to have the given address. This is really
	// nothing more than a hack, but it's very effective to force a public IP address
	// when one is known in advance for the media server. In doing so, KMS will not
	// need a STUN or TURN server, but remote peers will still be able to contact it.
	// </p>
	// <p>
	// You can try using this setting if KMS is deployed on a publicly accessible
	// server, without NAT, and with a static public IP address. But if it doesn't
	// work for you, just go back to configuring a STUN or TURN server for ICE.
	// </p>
	// <p>
	// Only set this parameter if you know what you're doing, and you understand 100%
	// WHY you need it. For the majority of cases, you should just prefer to
	// configure a STUN or TURN server.
	// </p>
	// <p><code>externalIPv6</code> is a single IPv6 address.</p>
	// <p>Example:</p>
	// <ul>
	// <li><code>externalIPv6=2001:0db8:85a3:0000:0000:8a2e:0370:7334</code></li>
	// </ul>
	//
	ExternalIPv6 string

	// External IP address of the media server.
	// <p>
	// Forces all local IPv4 and IPv6 ICE candidates to have the given address. This
	// is really nothing more than a hack, but it's very effective to force a public
	// IP address when one is known in advance for the media server. In doing so, KMS
	// will not need a STUN or TURN server, but remote peers will still be able to
	// contact it.
	// </p>
	// <p>
	// You can try using this setting if KMS is deployed on a publicly accessible
	// server, without NAT, and with a static public IP address. But if it doesn't
	// work for you, just go back to configuring a STUN or TURN server for ICE.
	// </p>
	// <p>
	// Only set this parameter if you know what you're doing, and you understand 100%
	// WHY you need it. For the majority of cases, you should just prefer to
	// configure a STUN or TURN server.
	// </p>
	// <p><code>externalAddress</code> is a single IPv4 or IPv6 address.</p>
	// <p>Examples:</p>
	// <ul>
	// <li><code>externalAddress=198.51.100.1</code></li>
	// <li><code>externalAddress=2001:0db8:85a3:0000:0000:8a2e:0370:7334</code></li>
	// </ul>
	// @deprecated Use <code>externalIPv4</code> and/or <code>externalIPv6</code> instead.
	//
	ExternalAddress string

	// the ICE candidate pair (local and remote candidates) used by the ICE library for each stream.
	ICECandidatePairs []*IceCandidatePair

	// the ICE connection state for all the connections.
	IceConnectionState []*IceConnection
}

type WebRtcEndpoint_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	MediaPipeline      core.MediaPipeline
	Recvonly           bool
	Sendonly           bool
	UseDataChannels    bool
	CertificateKeyType CertificateKeyType
	QosDscp            DSCPValue
}

func (WebRtcEndpoint_builder) GetTypeName() string {
	return "WebRtcEndpoint"
}

type WebRtcEndpointGatherCandidatesParams struct {
}

func (WebRtcEndpointGatherCandidatesParams) OperationName() string {
	return "gatherCandidates"
}

// Start the ICE candidate gathering.
// <p>
// This method triggers the asynchronous discovery of ICE candidates (as per the
// Trickle ICE mechanism), and returns immediately. Every newly trickled
// candidate is reported to the application by means of an
// <code>IceCandidateFound</code> event. Finally, when all candidates have been
// gathered, the <code>IceGatheringDone</code> event is emitted.
// </p>
// <p>
// Normally, you would call this method as soon as possible after calling
// <code>SdpEndpoint::generateOffer</code> or
// <code>SdpEndpoint::processOffer</code>, to quickly start discovering
// candidates and sending them to the remote peer.
// </p>
// <p>
// You can also call this method <em>before</em> calling
// <code>generateOffer</code> or <code>processOffer</code>. Doing so will include
// any already gathered candidates into the resulting SDP. You can leverage this
// behavior to implement fully traditional ICE (without Trickle): first call
// <code>gatherCandidates</code>, then only handle the SDP messages after the
// <code>IceGatheringDone</code> event has been received. This way, you're making
// sure that all candidates have indeed been gathered, so the resulting SDP will
// include all of them.
// </p>
func (elem *WebRtcEndpoint) GatherCandidates(ctx context.Context) error {
	request := kurento.BuildInvoke(elem.Id, &WebRtcEndpointGatherCandidatesParams{})

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type WebRtcEndpointAddIceCandidateParams struct {
	Candidate IceCandidate `json:"Candidate"`
}

func (WebRtcEndpointAddIceCandidateParams) OperationName() string {
	return "addIceCandidate"
}

// Process an ICE candidate sent by the remote peer of the connection.
func (elem *WebRtcEndpoint) AddIceCandidate(ctx context.Context, params *WebRtcEndpointAddIceCandidateParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type WebRtcEndpointCreateDataChannelParams struct {
	Label             string `json:"Label"`
	Ordered           bool   `json:"Ordered"`
	MaxPacketLifeTime int    `json:"MaxPacketLifeTime"`
	MaxRetransmits    int    `json:"MaxRetransmits"`
	Protocol          string `json:"Protocol"`
}

func (WebRtcEndpointCreateDataChannelParams) OperationName() string {
	return "createDataChannel"
}

// Create a new data channel, if data channels are supported.
// <p>
// Being supported means that the WebRtcEndpoint has been created with data
// channel support, the client also supports data channels, and they have been
// negotiated in the SDP exchange. Otherwise, the method throws an exception,
// indicating that the operation is not possible.
// </p>
// <p>
// Data channels can work in either unreliable mode (analogous to User Datagram
// Protocol or UDP) or reliable mode (analogous to Transmission Control Protocol
// or TCP). The two modes have a simple distinction:
// </p>
// <ul>
// <li>
// Reliable mode guarantees the transmission of messages and also the order in
// which they are delivered. This takes extra overhead, thus potentially making
// this mode slower.
// </li>
// <li>
// Unreliable mode does not guarantee every message will get to the other side
// nor what order they get there. This removes the overhead, allowing this mode
// to work much faster.
// </li>
// </ul>
// <p>If data channels are not supported, this method throws an exception.</p>
func (elem *WebRtcEndpoint) CreateDataChannel(ctx context.Context, params *WebRtcEndpointCreateDataChannelParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}

type WebRtcEndpointCloseDataChannelParams struct {
	ChannelId int `json:"ChannelId"`
}

func (WebRtcEndpointCloseDataChannelParams) OperationName() string {
	return "closeDataChannel"
}

// Closes an open data channel
func (elem *WebRtcEndpoint) CloseDataChannel(ctx context.Context, params *WebRtcEndpointCloseDataChannelParams) error {
	request := kurento.BuildInvoke(elem.Id, params)

	// Returns error or nil
	_, err := kurento.CallSimple[any](ctx, elem.GetConnection(), request)
	if err != nil {
		return fmt.Errorf("rpc error: %w", err)
	}
	return nil

}
