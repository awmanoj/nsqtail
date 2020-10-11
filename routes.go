package main

import "github.com/gorilla/mux"

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/", handleIndexRequest)
	router.HandleFunc("/nsqtail/{topic}", handleNSQTailRequest)
}