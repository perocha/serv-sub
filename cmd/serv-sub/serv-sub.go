package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

var sub = &common.Subscription{
	PubsubName: "orderpubsub",
	Topic:      "orders",
	Route:      "/orders",
}

//
// Main entry point
//
func main() {
	// read app-port passed through 'dapr run' command line
	// Refer to https://docs.dapr.io/reference/cli/dapr-run/
	// for dapr flags and their corresponding environment variables
	/*
		appPort, isSet := os.LookupEnv("APP_PORT")
			if !isSet {
				log.Fatalf("--app-port is not set. Re-run dapr run with -p or --app-port.")
			}
	*/
	appPort := "6001"
	log.Printf("Starting Dapr Subscriber on port %s", appPort)

	s := daprd.NewService(":" + appPort)
	log.Printf("Subscribing to topic %s", sub.Topic)
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	fmt.Println("Subscriber received: ", e.Data)
	return false, nil
}
