package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
type MiddleWareHandler struct{
	r *httprouter.Router
	l *ConnLimiter
}
func NewMiddleWareHandler(r *httprouter.Router,cc int)*MiddleWareHandler{
	middleWareHandler:=&MiddleWareHandler{}
	middleWareHandler.r=r
	middleWareHandler.l=NewConnLimiter(cc)
	return middleWareHandler
}
func (m MiddleWareHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if !m.l.GetConn(){
		sendErrorResponse(w,http.StatusTooManyRequests,"too many requests")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}
func RegisterHander()*httprouter.Router{
	router:=httprouter.New()
	router.GET("/video/:video_id",StreamHandler)
	router.POST("/upload/:video_id",UploadHandler)
	router.GET("/test",TestHandler)
	router.GET("/token",GetToken)
    return router
}
func main(){
    r:=RegisterHander()
    mh:=NewMiddleWareHandler(r,2)
    //fmt.Println("listen")
    http.ListenAndServe("127.0.0.1:9001",mh)
}
