package main

import "github.com/ndane/domainstatus/internal/nats"

func main() {
	nats.Connect("localhost:4222")
}
