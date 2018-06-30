package session

import (
	"sync"
	"MyProject/2017_6_23video_server/api/dbops"
	"MyProject/2017_6_23video_server/api/defs"
	"MyProject/2017_6_23video_server/api/utils"
	"time"
)


var sessionMap *sync.Map

func init(){
	sessionMap=&sync.Map{}
}
func nowInMilli()int64{
	return time.Now().UnixNano()/100000
}
func LoadSessionFromDB(){
    r,err:=dbops.GetAllSession()
    if err!=nil{
    	return
	}
	r.Range(func(k,v interface{})bool {
		ss:=v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})

}
func GenerateNewSessionId(un string)string{
	id,_:=utils.NewUUID()
	ct:=nowInMilli()
	ttl:=ct+30*60*1000
	ss:=&defs.SimpleSession{un,ttl}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id,ttl,un)
	return id
}

func IsSessionExpired(sid string)(string,bool){
    ss,ok:=sessionMap.Load(sid)
    if ok{
    	ct:=nowInMilli()
 		if ss.(*defs.SimpleSession).TTL<ct{
			DeleteExpiredSession(sid)
 			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}

func DeleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}