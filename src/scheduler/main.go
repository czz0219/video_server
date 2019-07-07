package main

import (
	"log"
	"net/http"
	"scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// restlet client 测试: http://127.0.0.1:9001/video-delete-record/123
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}
func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	log.Println("lensting 9001")
	http.ListenAndServe(":9001", r)

}
