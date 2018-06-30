package main

import (
	"net/http"
	"MyProject/2017_6_23video_server/api/session"
	"MyProject/2017_6_23video_server/api/defs"
)
var HEADER_FIELD_SESSION="X-Session-Id"
var HEADER_FIELD_UNAME="X-User-Name"

func ValidateUserSession(r *http.Request) bool{
	sid:=r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid)==0{
		return false
	}
	uname,ok:=session.IsSessionExpired(sid)
	if ok{
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME,uname)
	return true
}

func ValidateUser(r *http.Request,w http.ResponseWriter)bool{
	//uname:=r.Header.Get(HEADER_FIELD_UNAME)
	_,err:=r.Cookie("username")
	if err!=nil{
		sendErrorResponse(w,defs.ErrorUserNotAuth)
		return false
	}
	return true
}
