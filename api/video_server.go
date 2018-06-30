package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"MyProject/2017_6_23video_server/api/session"
	"siliang/config"
)
type MiddleWareHandler struct{
	r *httprouter.Router
}
func NewMiddleWareHandler(r *httprouter.Router)*MiddleWareHandler{
	middleWareHandler:=&MiddleWareHandler{}
	middleWareHandler.r=r
	return middleWareHandler
}
func (m MiddleWareHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	m.r.ServeHTTP(w,r)
}
func RegisterHandler() *httprouter.Router{
	router:=httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/:username",LoginUser)
	router.GET("/user/:username",GetUserInfo)
	router.POST("/user/:username/videos",AddNewVideo)
	router.GET("/user/:username/videos",ListAllVideos)
    router.DELETE("/user/:username/videos/:video_id",DeleteVideo)
	router.POST("/videos/:video_id/comments",PostComment)
	router.GET("/videos/:video_id/comments",ShowComments)

	return router
}
func Prepare(){
	session.LoadSessionFromDB()
}
func main(){
	Prepare()
    r:=RegisterHandler()
    m:=NewMiddleWareHandler(r)
    http.ListenAndServe(config.GetLBAddr()+":8088",m)
}
