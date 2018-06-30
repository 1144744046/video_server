package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func RegisterHandler() *httprouter.Router{
	router:=httprouter.New()
	router.GET("/",HomeHandler)
	router.POST("/",HomeHandler)
	router.GET("/userhome",UserHomeHandler)
	router.POST("/userhome",UserHomeHandler)
	router.POST("/api",ApiHandler)
	router.POST("/upload/:video_id",ProxyHandler)
	router.GET("/video/:video_id",ProxyHandler)
	router.GET("/token",ProxyHandler)
	router.ServeFiles("/static/*filepath",http.Dir("./template"))
	return router
}
func main(){
	r:=RegisterHandler()
	http.ListenAndServe(":8080",r)
}
