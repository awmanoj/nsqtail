package main

import (
	//"errors"
	"flag"
	"github.com/awmanoj/nsqtail/html"
	"github.com/awmanoj/nsqtail/nsq"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var NSQLookupdAddrPtr = flag.String("nsqlookupd", "127.0.0.1:4161", "NSQLookupd Address")

func main() {
	port := os.Getenv("PORT")

	// global configuration captured via command line parameter
	os.Setenv(nsq.LookupdAddrEnv, *NSQLookupdAddrPtr)

	nsq.InitConsumers()

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndexRequest)
	r.HandleFunc("/nsqtail/{topic}", handleNSQTailRequest)

	// yaay!! start the server!
	log.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

// Index page, show all topics as links
func handleIndexRequest(w http.ResponseWriter, r *http.Request) {
	topics, err := nsq.GetTopics()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("err", "problem fetching topics %v\n", err)
		return
	}

	html.IndexHTMLTemplate.Execute(w, html.IndexHTMLData{
		NSQLookupdAddress: os.Getenv(nsq.LookupdAddrEnv),
		Topics: topics.Topics,
	})
}

// fetch last 10 messages on the topic
func handleNSQTailRequest(w http.ResponseWriter, r *http.Request) {
	topic := mux.Vars(r)["topic"]

	lastNRequests, err := nsq.FetchLastNRequests(topic, nsq.MaxNumOfMessages)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("err", "problem fetching last %d requests %v\n", nsq.MaxNumOfMessages, err)
		return
	}

	html.TailHTMLTemplate.Execute(w, html.TailHTMLData{
		NSQLookupdAddress: os.Getenv(nsq.LookupdAddrEnv),
		Topic: topic,
		MessageCount: len(lastNRequests),
		Messages: lastNRequests,
	})
}
