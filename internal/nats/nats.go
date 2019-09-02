package nats

import (
	"encoding/json"
	"fmt"
	"sync"

	nats "github.com/nats-io/nats.go"
	"github.com/ndane/domainstatus/pkg/messages"
)

// Connect - Setup connection to nats server
func Connect(address string) {
	nc, _ := nats.Connect(address)

	hb := &messages.Heartbeat{
		Address: "localhost:8080",
	}

	b, _ := json.Marshal(hb)
	nc.Publish("hb", b)

	nc.Subscribe("hb", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
