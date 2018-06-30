package dbops

import (
	"database/sql"
	"MyProject/2017_6_23video_server/api/defs"
	"MyProject/2017_6_23video_server/api/utils"
	"time"
	"fmt"
)

func AddUser(username string,password string)error{
   stmIns,err:=sqlconn.Prepare("insert into users (user_name,pwd) values (?,?)")
   if err!=nil{
   	 fmt.Println(err)
   	 return err
   }
   _,err=stmIns.Exec(username,password)
	if err!=nil{
		fmt.Println(err)
		return err
	}
   defer stmIns.Close()
   return nil
}

func GetUser(username string)(*defs.User,error){
	stmIns,err:=sqlconn.Prepare("select id,pwd from users where user_name=?")
	if err!=nil{
		return nil,err
	}
	var pwd string
	var id int
	err=stmIns.QueryRow(username).Scan(&id,&pwd)
	if err!=nil && err!=sql.ErrNoRows{
		return nil,err
	}
	defer stmIns.Close()

	return &defs.User{id,username,pwd},nil
}
func DeleteUser(username string,password string)error{
	stmIns,err:=sqlconn.Prepare("delete from users where user_name=? and pwd=?")
	if err!=nil{
		return err
	}
	_,err=stmIns.Exec(username,password)
	if err!=nil{
		return err
	}
	defer stmIns.Close()
	return nil
}

func AddNewVidio(aid int,name string)(*defs.VidioInfo,error){
	vid,err:=utils.NewUUID()
	if err!=nil{
		return nil,err
	}
	t:=time.Now()
	ctime:=t.Format("2006-2-16 15:04:05")
	stmIns,err:=sqlconn.Prepare("insert into vidio_info (id,author_id,name,display_ctime) values(?,?,?,?)")
	if err!=nil{
		return nil,err
	}
	_,err=stmIns.Exec(vid,aid,name,ctime)
	if err!=nil{
		return nil,err
	}
	defer stmIns.Close()
	vidio_info:=&defs.VidioInfo{Id:vid,Authid:aid,Name:name,DisplayCtime:ctime}
	return vidio_info,nil
}

func GetVidioInfo(vid string)(*defs.VidioInfo,error){
	stmIns,err:=sqlconn.Prepare("select author_id,name,display_ctime from vidio_info where id=?")
	if err!=nil{
		return nil,err
	}
	var author_id int
	var name string
	var display_ctime string
	err=stmIns.QueryRow(vid).Scan(&author_id,&name,&display_ctime)
	if err!=nil && err!=sql.ErrNoRows{
		return nil,err
	}
	vidio_info:=&defs.VidioInfo{Id:vid,Authid:author_id,Name:name,DisplayCtime:display_ctime}
	defer stmIns.Close()
	return vidio_info,nil
}

func DeleteVidioInfo(vid string)error{
	stmIns,err:=sqlconn.Prepare("delete from vidio_info where id=?")
	if err!=nil{
		return err
	}
	_,err=stmIns.Exec(vid)
	if err!=nil{
		return err
	}
	defer stmIns.Close()
	return nil
}

func AddNewComment(vid string,aid int,content string)error{
	cid,err:=utils.NewUUID()
	if err!=nil{
		return err
	}
	stmIns,err:=sqlconn.Prepare("insert into comment (id,vidio_id,author_id,content) values(?,?,?,?)")
	if err!=nil{
		return err
	}
	stmIns.Exec(cid,vid,aid,content)
	defer stmIns.Close()
	return nil
}

func ListComments(vid string)([]*defs.Comment,error){
	stmIns,err:=sqlconn.Prepare(`select comment.id,users.user_name,comment.content,comment.author_id
from comment inner join user on comment.author_id=users.id
where comment.vidio_id=?`)
	if err!=nil{
		return nil,err
	}
	rows,err:=stmIns.Query(vid)
	var res []*defs.Comment
	if err!=nil{
		return res,err
	}
	for rows.Next(){
       var id,name,content string
       var author_id int
       if err=rows.Scan(&id,&name,&content,&author_id);err!=nil{
       	    return res,err
	   }
	   c:=&defs.Comment{id,vid,name,content,author_id}
       res=append(res,c)
	}
	defer stmIns.Close()
	return res,nil
}
func ListVideoInfo(un string)([]*defs.VidioInfo,error){
	stmIns,err:=sqlconn.Prepare(`select vidio_info.id,vidio_info.author_id,vidio_info.name,display_ctime from vidio_info
inner join users on vidio_info.author_id=users.id and users.user_name=?`)
	if err!=nil{
		return nil,err
	}
	rows,err:=stmIns.Query(un)
	var res []*defs.VidioInfo
	if err!=nil{
		return res,err
	}
	for rows.Next(){
		var id,name,display_ctime string
		var author_id int
		if err=rows.Scan(&id,&author_id,&name,&display_ctime);err!=nil{
			return res,err
		}
		c:=&defs.VidioInfo{id,author_id,name,display_ctime}
		res=append(res,c)
	}
	defer stmIns.Close()
	return res,nil
}