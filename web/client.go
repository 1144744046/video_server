package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"io"
	"encoding/json"
	"bytes"
	"net/url"
	"siliang/config"
)
var httpclient *http.Client

func init(){
	httpclient=&http.Client{}
}
func request(body *ApiBody,w http.ResponseWriter,r *http.Request){
	var resp *http.Response
	var err error
	u,_:=url.Parse(body.Url)
    u.Host=config.GetLBAddr()+":"+u.Port()
    url:=u.String()
	switch body.Method {
	case http.MethodGet:
		req,_:=http.NewRequest("GET",url,nil)
		req.Header=r.Header
		resp,err=httpclient.Do(req)
		if err!=nil{
			log.Println(err)
			return
		}
		normalResponse(w,resp)
	case http.MethodPost:
		req,_:=http.NewRequest("POST",url,bytes.NewBuffer([]byte(body.ReqBody)))
		req.Header=r.Header
		resp,err=httpclient.Do(req)
		if err!=nil{
			log.Println(err)
			return
		}
		normalResponse(w,resp)
	case http.MethodDelete:
		req,_:=http.NewRequest("DELETE",url,nil)
		req.Header=r.Header
		resp,err=httpclient.Do(req)
		if err!=nil{
			log.Println(err)
			return
		}
		normalResponse(w,resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w,"bad request")
	}
}

func normalResponse(w http.ResponseWriter,r *http.Response){
    res,err:=ioutil.ReadAll(r.Body)
    if err!=nil{
		re,_:=json.Marshal(ErrorResponseBodyParseFail)
		w.WriteHeader(500)
		io.WriteString(w,string(re))
		return
	}
	w.WriteHeader(r.StatusCode)
	io.WriteString(w,string(res))
}