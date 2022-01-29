package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	dapr "github.com/dapr/go-sdk/client"
)

var pubsubName = "pubsub"
var ctx = context.Background()

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/addVideo", handlerAddVideo)
	http.HandleFunc("/addPodcast", handlerAddPodcast)
	http.HandleFunc("/addBlog", handlerAddBlog)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	log.Println("Starting the server on " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	output := "This is a silly demo"
	output = fmt.Sprintf("%s\n", output)
	fmt.Fprintf(w, output)
}

func handlerAddVideo(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	url := req.URL.Query().Get("url")
	data := fmt.Sprintf(`id: %s
name: %s
url: %s`,
		id,
		name,
		url,
	)
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	if err := client.PublishEvent(ctx, pubsubName, "video", []byte(data)); err != nil {
		panic(err)
	}
	output := fmt.Sprintf("Added video:\n%s\n", data)
	log.Println(output)
	fmt.Fprintf(w, output)
}

func handlerAddPodcast(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	url := req.URL.Query().Get("url")
	data := fmt.Sprintf(`id: %s
name: %s
url: %s`,
		id,
		name,
		url,
	)
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	if err := client.PublishEvent(ctx, pubsubName, "podcast", []byte(data)); err != nil {
		panic(err)
	}
	output := fmt.Sprintf("Added podcast:\n%s\n", data)
	fmt.Fprintf(w, output)
}

func handlerAddBlog(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	url := req.URL.Query().Get("url")
	data := fmt.Sprintf(`name: %s
url: %s`,
		name,
		url,
	)
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	if err := client.PublishEvent(ctx, pubsubName, "blog", []byte(data)); err != nil {
		panic(err)
	}
	output := fmt.Sprintf("Added podcast:\n%s\n", data)
	fmt.Fprintf(w, output)
}
