package dbops

import _ "github.com/go-sql-driver/mysql"
func ReadVideoDeleteRecord(count int)([]string,error){
	stmIns,err:=sqlconn.Prepare("select video_id from vidio_del_rec limit ?")
	if err!=nil{
		return nil,err
	}
    var ids []string
    rows,err:=stmIns.Query(count)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var id string
		if err:=rows.Scan(&id);err!=nil{
			return nil,err
		}
		ids=append(ids,id)
	}
	defer stmIns.Close()
	return ids,nil
}

func DelVideoRecord(vid string)error{
	stmIns,err:=sqlconn.Prepare("delete from vidio_del_rec where video_id=?")
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