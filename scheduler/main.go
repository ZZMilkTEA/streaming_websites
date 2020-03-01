package main

import (
	"Streaming_websites/scheduler/taskRunner"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}
func main() {
	log.Printf("Scheduler start!")
	go taskRunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":13112", r)
}
