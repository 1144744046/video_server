package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"MyProject/2017_6_23video_server/scheduler/task"
	"siliang/config"
)

func RegisterHander()*httprouter.Router{
	router:=httprouter.New()
	router.GET("/video_delete_record/:vidio_id",VideoDeleteHandler)

	return router
}

func main(){
	go task.Start()
   r:=RegisterHander()
   http.ListenAndServe(config.GetLBAddr()+":9002",r)
}
