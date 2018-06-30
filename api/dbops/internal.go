package dbops

import (
	"strconv"
	"MyProject/2017_6_23video_server/api/defs"
	"database/sql"
	"sync"
)

func InsertSession(sid string,ttl int64,user_name string) error{
	ttlstr:=strconv.FormatInt(ttl,10)
	stmIns,err:=sqlconn.Prepare("insert into session (session_id,login_name,ttl)values(?,?,?)")
	if err!=nil{
		return err
	}
	_,err=stmIns.Exec(sid,user_name,ttlstr)
	if err!=nil{
		return err
	}
	defer stmIns.Close()
	return nil
}

func GetSession(sid string)(*defs.SimpleSession,error){
	 ss:=&defs.SimpleSession{}
	 stmIns,err:=sqlconn.Prepare("select login_name,ttl from session where session_id=?")
	if err!=nil{
		return nil,err
	}
	var login_name string
	var ttl string
	err=stmIns.QueryRow(sid).Scan(&login_name,&ttl)
	if err!=nil && err!=sql.ErrNoRows{
		return nil,err
	}
	ss.Username=login_name
	ss.TTL,err=strconv.ParseInt(ttl,10,64)
	if err!=nil{
		return nil,err
	}
	defer stmIns.Close()
	return ss,nil
}

func GetAllSession()(*sync.Map,error){
	stmIns,err:=sqlconn.Prepare(`select session_id,login_name,ttl from session`)

	if err!=nil{
		return nil,err
	}
	rows,err:=stmIns.Query()
	res:=&sync.Map{}
	if err!=nil{
		return res,err
	}
	for rows.Next(){
		var session_id,login_name,ttl string
		if err=rows.Scan(&session_id,&login_name,&ttl);err!=nil{
			break
		}
		ttl_int,err:=strconv.ParseInt(ttl,10,64)
		if err!=nil{
			continue
		}
		ss:=&defs.SimpleSession{login_name,ttl_int}
		res.Store(session_id,ss)
	}
	defer stmIns.Close()
	return res,nil
}

func DeleteSession(sid string)error{
	stmIns,err:=sqlconn.Prepare("delete from session where session_id=?")
	if err!=nil{
		return err
	}
	_,err=stmIns.Exec(sid)
	if err!=nil{
		return err
	}
	defer stmIns.Close()
	return nil
}