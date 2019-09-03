package main

import (
	"sync"

	"github.com/ndane/domainstatus/pkg/heartbeat"
)

func main() {
	heartbeat.ConnectAndStart("domainstatus", "localhost:4222")
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
