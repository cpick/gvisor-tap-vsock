package types

type TransportProtocol string

type TransportProtocols map[TransportProtocol]struct{}

const (
	UDP   TransportProtocol = "udp"
	TCP   TransportProtocol = "tcp"
	TCPFD TransportProtocol = "tcpfd"
	UNIX  TransportProtocol = "unix"
	NPIPE TransportProtocol = "npipe"
)

func DefaultTransportProtocols() TransportProtocols {
	return TransportProtocols{
		UDP:   struct{}{},
		TCP:   struct{}{},
		UNIX:  struct{}{},
		NPIPE: struct{}{},
	}
}

type ExposeRequest struct {
	Local    string            `json:"local"`
	Remote   string            `json:"remote"`
	Protocol TransportProtocol `json:"protocol"`
}

type UnexposeRequest struct {
	Local    string            `json:"local"`
	Protocol TransportProtocol `json:"protocol"`
}
