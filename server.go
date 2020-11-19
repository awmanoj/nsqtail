package main

import (
	//"errors"
	"flag"
	"github.com/awmanoj/nsqtail/nsq"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var NSQLookupdAddrPtr = flag.String("nsqlookupd", "127.0.0.1:4161", "NSQLookupd Address")

func main() {
	flag.Parse()

	log.Println("===", os.Args)
	// global configuration captured via command line parameter
	os.Setenv(nsq.LookupdAddrEnv, *NSQLookupdAddrPtr)

	nsq.InitConsumers()

	router := mux.NewRouter()
	InitRoutes(router)

	// yaay!! start the server!
	log.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
