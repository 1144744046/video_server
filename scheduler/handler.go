package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"MyProject/2017_6_23video_server/scheduler/dbops"
)

func VideoDeleteHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
     vid:=p.ByName("vidio_id")
	if len(vid)==0{
		sendErrorResponse(w,http.StatusBadRequest,"get vid null")
		return
	}
	 err:=dbops.AddVideoRecord(vid)
	 if err!=nil{
		 sendErrorResponse(w,http.StatusInternalServerError,"db ops err")
		 return
	 }
	 sendErrorResponse(w,200,"success")
	 return
}
