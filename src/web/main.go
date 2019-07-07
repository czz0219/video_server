package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {

	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	router.POST("/upload/:vid-id",proxyHandler)
	router.GET("/videos/:vid-id",proxyHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	//127.0.0.1:8000/statics/ == ./template
	//template自动挂载到 statics下
	return router
}
func main() {
	r := RegisterHandler()
	log.Println("lensting 8080")
	http.ListenAndServe(":8080", r)

}
