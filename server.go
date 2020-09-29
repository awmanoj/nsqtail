package main

import (
	//"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

	// extract the query parameters
	query := r.URL.Query()

	// list of query parameters with key 'n'
	ns, ok := query["n"]
	if !ok || len(ns) == 0 {
		// if nothing else, assume ?n=10
		ns = append(ns, "10")
	}

	// flag to track if continues updates needed (f=true)
	var f bool
	// with err handling above, it is guaranteed to have at least one value
	l := len(ns[0])
	if l != 0 {
		// look for 'f' flag only at the end of the parameter value
		// ?n=100f is valid ?n=100fe is not
		f = ns[0][l-1] == 'f'
	}

	n, err := strconv.Atoi(ns[0])
	if err != nil {
		n = 10
	}

	fmt.Fprintf(w, "%s -- All is well!\n [%d] [%v]", topic, n, f)
}
