package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"

	"io"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"fmt"
	"net/http/httputil"
	"siliang/config"
)
type HomePage struct{
	Name string
}
type UserPage struct{
	Name string
}
func HomeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
   cname,err1:=r.Cookie("username")
   sid,err2:=r.Cookie("session")
   if err1!=nil ||err2!=nil{
	   page:=&HomePage{"siliang"}
	   t,e:=template.ParseFiles("./template/home.html")
	   if e!=nil{
		   return
	   }
	   t.Execute(w,page)
	   return
   }
   if len(cname.Value)!=0 && len(sid.Value)!=0{
   	  http.Redirect(w,r,"/userhome",http.StatusFound)
   	  return
   }

}

func UserHomeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	cname,err1:=r.Cookie("username")
	_,err2:=r.Cookie("session")
	if err1!=nil ||err2!=nil{
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}

	fname:=r.FormValue("username")
	var page *UserPage
	if len(cname.Value)!=0{
		page=&UserPage{cname.Value}
	}else if len(fname)!=0{
		page=&UserPage{fname}
	}
	t,e:=template.ParseFiles("./template/userhome.html")
    if e!=nil{
    	return
	}
    t.Execute(w,page)
}

func ApiHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
    if r.Method!=http.MethodPost{
    	re,_:=json.Marshal(ErrorBadRequest)
    	io.WriteString(w,string(re))
		return
	}
    res,_:=ioutil.ReadAll(r.Body)
    apibody:=&ApiBody{}
    fmt.Println(string(res))
    if err:=json.Unmarshal(res,apibody);err!=nil{
    	fmt.Println("序列化出错拉",err)
		re,_:=json.Marshal(ErrorBadRequest)
		io.WriteString(w,string(re))
		return
	}
    request(apibody,w,r)
   defer r.Body.Close()
}

func ProxyHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	u,_:=url.Parse("http://"+config.GetLBAddr()+":9001/")
	fmt.Println(u.String())
	proxy:=httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w,r)
}

