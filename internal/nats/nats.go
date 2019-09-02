package nats

import nats "github.com/nats-io/nats.go"

// Conenct - Setup connection to nats server
func Connect() {
	nc, _ := nats.Connect("")

}
