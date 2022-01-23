package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
)

var ctx = context.Background()
var data = []byte("hello")
var store = "statestore" // TODO: Rewrite

func main() {
	log.Println("Starting the server on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	subscriber()
}

func subscriber() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	item, err := client.GetState(ctx, store, "A")
	if err != nil {
		panic(err)
	}
	fmt.Printf("data [key: %s etag: %s]: %s", item.Key, item.Etag, string(item.Value))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is a silly demo")
}
