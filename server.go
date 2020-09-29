package main

import (
	//"errors"
	"fmt"
	"log"
	"net/http"
	"os"

  "github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/nsqtail/{topic}", handleNSQTailRequest)

	// yaay!! start the server!
	log.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func handleNSQTailRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// extract sunsign from variable
	vars := mux.Vars(r)
	topic := vars["topic"]

	fmt.Fprintf(w, topic + " -- All is well!")
}

