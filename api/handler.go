package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"MyProject/2017_6_23video_server/api/defs"
	"encoding/json"
	"MyProject/2017_6_23video_server/api/dbops"
	"MyProject/2017_6_23video_server/api/session"
	"log"
	"fmt"
)

func CreateUser(w http.ResponseWriter,r *http.Request,param httprouter.Params){
   // io.WriteString(w,"http create user handler")
    res,_:=ioutil.ReadAll(r.Body)
    user:=&defs.CreateUser{}
    fmt.Println(string(res))
    if err:=json.Unmarshal(res,user);err!=nil{
    	fmt.Println(err)
    	sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
    	return
	}
    if err:=dbops.AddUser(user.Username, user.Pwd);err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	//如果创建成功，返回session
	sid:=session.GenerateNewSessionId(user.Username)
	su:=&defs.SignedUp{true,sid}
    if resp,err:=json.Marshal(su);err!=nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}else{
		fmt.Println(string(resp))
		sendNormalResponse(w,string(resp),201)
	}
}

func LoginUser(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	res,_:=ioutil.ReadAll(r.Body)
	ubody:=&defs.CreateUser{}
	if err:=json.Unmarshal(res,ubody);err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	uname:=param.ByName("username")

    user,err:=dbops.GetUser(uname)
    if err!=nil || len(user.Pwd)==0 ||user.Pwd!=ubody.Pwd{
    	sendErrorResponse(w,defs.ErrorUserNotAuth)
    	return
	}
	id:=session.GenerateNewSessionId(uname)
	su:=defs.SignedUp{true,id}
	if suStr,err:=json.Marshal(su);err!=nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else{
		sendNormalResponse(w,string(suStr),200)
	}
	//io.WriteString(w,uname)
}

func GetUserInfo(w http.ResponseWriter,r *http.Request,param httprouter.Params){
    if !ValidateUser(r,w){
    	log.Println("not auth user")
    	return
	}
	uname:=param.ByName("username")
	user,err:=dbops.GetUser(uname)
	if err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
    //u:=&defs.User{user.Id,uname,user.Pwd}
    if resp,err:=json.Marshal(user);err!=nil{
    	sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
    	return
	}else{
		sendNormalResponse(w,string(resp),200)
	}

}
func AddNewVideo(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	if !ValidateUser(r,w){
		log.Println("not auth user")
		return
	}
	res,_:=ioutil.ReadAll(r.Body)
	vbody:=&defs.VidioInfo{}
	if err:=json.Unmarshal(res,vbody);err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	uname:=param.ByName("username")
	user,err:=dbops.GetUser(uname)
	vi,err:=dbops.AddNewVidio(user.Id,vbody.Name)
	if err!=nil {
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}
	if resp,err:=json.Marshal(vi);err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}else{
		sendNormalResponse(w,string(resp),201)
	}

}
func ListAllVideos(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	if !ValidateUser(r,w){
		log.Println("not auth user")
		return
	}
	uname:=param.ByName("username")
	vs,err:=dbops.ListVideoInfo(uname)
	if err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	if resp,err:=json.Marshal(vs);err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}else{
		sendNormalResponse(w,string(resp),200)
	}

}
func DeleteVideo(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	if !ValidateUser(r,w){
		log.Println("not auth user")
		return
	}
	vid:=param.ByName("video_id")
	err:=dbops.DeleteVidioInfo(vid)
	if err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	sendNormalResponse(w,"",204)
}
func PostComment(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	if !ValidateUser(r,w){
		log.Println("not auth user")
		return
	}
	res,_:=ioutil.ReadAll(r.Body)
	cbody:=&defs.Comment{}
	if err:=json.Unmarshal(res,cbody);err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	vid:=param.ByName("video_id")
	//dbops.GetUser()
	if err:=dbops.AddNewComment(vid,cbody.Author_id,cbody.Content);err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}else{
		sendNormalResponse(w,"ok",201)
	}

}
func ShowComments(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	if !ValidateUser(r,w){
		log.Println("not auth user")
		return
	}
	vid:=param.ByName("video_id")
	cm,err:=dbops.ListComments(vid)
	if err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	if resp,err:=json.Marshal(cm);err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}else{
		sendNormalResponse(w,string(resp),200)
	}
}
