package messages

const (
	// HeartbeatSubjectDomain - Hearbeat message subject
	HeartbeatSubjectDomain = "hb"
)

// Heartbeat - Message received from a hearbeat
type Heartbeat struct {
	Address              string  `json:"addr"`
	ServiceType          string  `json:"svc"`
	Load1                float64 `json:"load1"`
	Load5                float64 `json:"load5"`
	Load15               float64 `json:"load15"`
	MemoryUsedPercentage float64 `json:"memu"`
	MemoryTotal          uint64  `json:"memt"`
}
