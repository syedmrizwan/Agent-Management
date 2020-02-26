package main

import (
	"encoding/json"
	"log"
	stan "github.com/nats-io/stan.go"
)

type AgentConfig struct {
	DeviceHost string `json:"device_host"`
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`
}

func run() error {
	clusterID := "test-cluster"
	clientID := "manager"
	channelSubject := "agent.1AX24"
	conn, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL("127.0.0.1:4222"),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	agentConfig := &AgentConfig{
		DeviceHost: "192.16.1.2",
		DeviceID:   "1A2HUz",
		DeviceName: "router",
	}

	startEventStruct := struct {
		Config    AgentConfig `json:"config"`
		EventType string      `json:"event_type"`
	}{
		*agentConfig,
		"start",
	}
	b, _ := json.Marshal(startEventStruct)
	error1 := conn.Publish(channelSubject, b)
	if error1 != nil {
		return error1
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}