package main

import (
	"fmt"
	"log"
	"time"
	"github.com/nats-io/nats.go"
)

func main() {
	opts := []nats.Option{nats.Name("NATS Request")}

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	channelSubject := "agent.1AX24.health"
	payload := "Healthy???"

	msg, err := nc.Request(channelSubject, []byte(payload), 2*time.Second)
	if err != nil {
		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		fmt.Println("DEAD")
	} else {
		log.Printf("Published [%s] : '%s'", channelSubject, payload)
		log.Printf("Recieved [%s] : '%s'", channelSubject, string(msg.Data))
	}

	nc.Flush()

}
