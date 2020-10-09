package main

import (
	//"errors"
	"flag"
	"fmt"
	"github.com/awmanoj/nsqtail/nsq"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var NSQLookupdAddrPtr = flag.String("nsqlookupd", "127.0.0.1:4161", "NSQLookupd Address")

func main() {
	port := os.Getenv("PORT")

	// global configuration captured via command line parameter
	os.Setenv(nsq.LookupdAddrEnv, *NSQLookupdAddrPtr)

	r := mux.NewRouter()
	r.HandleFunc("/nsqtail/{topic}", handleNSQTailRequest)

	// yaay!! start the server!
	log.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

// fetch last 10 messages on the topic
func handleNSQTailRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	n := nsq.MaxNumOfMessages

	topic := mux.Vars(r)["topic"]

	lastNRequests, err := nsq.FetchLastNRequests(topic, n)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("err", "problem fetching last %d requests %v\n", n, err)
		return
	}

	fmt.Fprintf(w, strings.Join(lastNRequests, "\n"))
}
