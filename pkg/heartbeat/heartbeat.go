package heartbeat

import (
	"encoding/json"
	"strings"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/ndane/domainstatus/pkg/messages"
	log "github.com/sirupsen/logrus"
)

var stop chan bool

// ConnectAndStart - Connects to the NATS server and starts broadcasting a heartbeat
func ConnectAndStart(serviceName, servers string) {
	connection, err := nats.Connect(servers)
	if err != nil {
		log.WithError(err).Panic("Failed to connect to NATS")
	}

	hbCount := 0
	posthb(&hbCount, serviceName, connection)
	ticker := time.NewTicker(2 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				posthb(&hbCount, serviceName, connection)
			case <-stop:
				return
			}
		}
	}()
}

// Stop - Stop broadcasting heartbeat
func Stop() {
	stop <- true
}

func posthb(counter *int, service string, conn *nats.Conn) {
	*counter = *counter + 1
	subject := strings.Join([]string{
		messages.HeartbeatSubjectDomain,
		service,
		string(*counter),
	}, ".")

	message := &messages.Heartbeat{
		Address:     "localhost:8080",
		ServiceType: service,
	}

	payload, err := json.Marshal(&message)
	if err != nil {
		log.WithError(err).Panic()
	}

	conn.Publish(subject, payload)
}
