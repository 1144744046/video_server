package main

import (
	"net/http"
	"io"
)

func sendErrorResponse(w http.ResponseWriter,rc int,errMsg string){
	w.WriteHeader(rc)
	io.WriteString(w,errMsg)
}
