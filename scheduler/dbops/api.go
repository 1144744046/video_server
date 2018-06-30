package dbops

import ("fmt"
	_ "github.com/go-sql-driver/mysql"
)

func AddVideoRecord(vid string)error{
	stmIns,err:=sqlconn.Prepare("insert into vidio_del_rec (video_id) values (?)")
	if err!=nil{
		return err
	}
	_,err=stmIns.Exec(vid)
	fmt.Println(vid)
	if err!=nil{
		return err
	}
	defer stmIns.Close()
	return nil
}
