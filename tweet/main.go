package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

var ctx = context.Background()
var data = []byte("hello")
var store = "statestore"

func main() {
	log.Println("Starting the server on 8080...")
	http.HandleFunc("/", rootHandler)
	s := daprd.NewService(":8080")
	topics := strings.Split(os.Getenv("TOPICS"), ",")
	for _, topic := range topics {
		subscription := &common.Subscription{
			PubsubName: os.Getenv("PUBSUB_NAME"),
			Topic:      topic,
			Route:      os.Getenv("ROUTE"),
		}
		if err := s.AddTopicEventHandler(subscription, eventHandler); err != nil {
			log.Fatalf("error adding topic subscription: %v", err)
		}
	}
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is a silly demo")
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	log.Println("Sending a Tweet")
	return false, nil
}
