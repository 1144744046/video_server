package main

import (
	"net/http"
	"MyProject/2017_6_23video_server/api/defs"
	"github.com/gin-gonic/gin/json"
	"io"
)

func sendErrorResponse(w http.ResponseWriter,errResponse defs.ErrorResponse){
	w.WriteHeader(errResponse.HttpRC)
	resStr,err:=json.Marshal(errResponse.Err)
	if err!=nil{
		return
	}
	io.WriteString(w,string(resStr))
}
func sendNormalResponse(w http.ResponseWriter,resp string,sc int){
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}