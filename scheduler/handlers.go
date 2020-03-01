package main

import (
	"Streaming_websites/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	log.Printf("video: %s will be deleted", vid)
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal sever error")
		return
	}

	sendResponse(w, 200, "")
	return
}
