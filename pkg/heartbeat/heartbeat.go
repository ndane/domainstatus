package heartbeat

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"

	"github.com/ndane/domainstatus/pkg/messages"
)

var stop chan bool

// ConnectAndStart - Connects to the NATS server and starts broadcasting a heartbeat
func ConnectAndStart(serviceName string, servers []string) {
	connection, err := nats.Connect(strings.Join(servers, ","))
	if err != nil {
		log.WithError(err).Error("Failed to connect to NATS")
		return
	}

	hbCount := 1
	posthb(hbCount, serviceName, connection)
	ticker := time.NewTicker(2 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				hbCount++
				posthb(hbCount, serviceName, connection)
			case <-stop:
				connection.Close()
				return
			}
		}
	}()
}

// Stop - Stop broadcasting heartbeat
func Stop() {
	stop <- true
}

func posthb(msgid int, service string, conn *nats.Conn) {
	vmem, err := mem.VirtualMemory()
	avg, err := load.Avg()
	clientID, err := conn.GetClientID()

	subject := strings.Join([]string{
		messages.HeartbeatSubjectDomain,
		service,
		strconv.FormatUint(clientID, 10),
		strconv.Itoa(msgid),
	}, ".")

	message := &messages.Heartbeat{
		Address:              "localhost:8080",
		ClientID:             clientID,
		ServiceType:          service,
		Load1:                avg.Load1,
		Load5:                avg.Load5,
		Load15:               avg.Load15,
		MemoryUsedPercentage: vmem.UsedPercent,
		MemoryTotal:          vmem.Total,
	}

	payload, err := json.Marshal(&message)
	if err != nil {
		log.WithError(err).Error()
	}

	conn.Publish(subject, payload)
}
