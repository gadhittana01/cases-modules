package utils

import (
	"log"

	pusher "github.com/pusher/pusher-http-go/v5"
)


func NewPusherClient(config *Config) *pusher.Client {
	if config.PusherAppID == "" || config.PusherKey == "" || config.PusherSecret == "" {
		log.Println("Warning: Pusher credentials not configured, Pusher events will be disabled")
		return nil
	}

	client := pusher.Client{
		AppID:   config.PusherAppID,
		Key:     config.PusherKey,
		Secret:  config.PusherSecret,
		Cluster: config.PusherCluster,
		Secure:  true,
	}

	return &client
}


func EmitPaymentStatus(client *pusher.Client, channel, event string, data interface{}) error {
	if client == nil {
		return nil
	}

	err := client.Trigger(channel, event, data)
	if err != nil {
		log.Printf("Failed to emit Pusher event: %v", err)
		return err
	}

	return nil
}
