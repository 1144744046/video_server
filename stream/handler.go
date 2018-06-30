package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
	"fmt"
	"io/ioutil"
	"io"
	"html/template"
	"github.com/astaxie/beego/logs"
)


func TestHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	tmp,_:=template.ParseFiles("video/testupload.html")
	tmp.Execute(w,nil)
}
func StreamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	fmt.Println("看视频拉")
    vid:=p.ByName("video_id")
    vl:=VIDIO_DIR+vid
    fmt.Println(vl)
    video,err:=os.Open(vl)

    if err!=nil{
    	sendErrorResponse(w,http.StatusInternalServerError,"open video fail")
    	return
	}
	w.Header().Set("Content-Type","video/mp4")
    http.ServeContent(w,r,"",time.Now(),video)
    defer video.Close()
}
func UploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
    r.Body=http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
    fmt.Println("upload")
    token:=r.Header.Get("token")
    if !checkToken(token){
    	logs.Info("upload multiput")
		sendErrorResponse(w,http.StatusInternalServerError,"Repeated submission")
    	return
	}
    if err:=r.ParseMultipartForm(MAX_UPLOAD_SIZE);err!=nil{
    	sendErrorResponse(w,http.StatusBadRequest,"file is too big")
    	return
	}
	file,_,err:=r.FormFile("file")
	if err!=nil{
		sendErrorResponse(w,http.StatusInternalServerError,"from file err")
		return
	}
	data,err:=ioutil.ReadAll(file)
	if err!=nil{
		sendErrorResponse(w,http.StatusInternalServerError,"file read err")
		return
	}

	fn:=p.ByName("video_id")
	err=ioutil.WriteFile(VIDIO_DIR+fn,data,0666)
	if err!=nil{
		sendErrorResponse(w,http.StatusInternalServerError,"write file err")
		return
	}
    deleteToken(token)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,fn)

}
func GetToken(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	w.WriteHeader(http.StatusOK)
	io.WriteString(w,getToken())
}

