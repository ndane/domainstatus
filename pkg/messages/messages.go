package messages

const (
	// HeartbeatSubjectDomain - Hearbeat message subject
	HeartbeatSubjectDomain = "hb"
)

// Heartbeat - Message received from a hearbeat
type Heartbeat struct {
	Address     string `json:"addr"`
	ServiceType string `json:"svc"`
	AverageLoad struct {
		// CPUPercentage - Average cpu load
		CPUPercentage float32 `json:"cpu"`

		// MemUsage - Current memory usage in %
		MemoryUsage float32 `json:"mem"`
	} `json:"load"`
}
