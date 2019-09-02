package messages

// MessageType - Denotes the type of message
type MessageType uint

const (
	// HeartbeatMessage - Hearbeat message
	HeartbeatMessage MessageType = iota
)

// Heartbeat - Message received from a hearbeat
type Heartbeat struct {
	Address string
}
