package task

import (
	"MyProject/2017_6_23video_server/scheduler/dbops"
	"errors"
	"os"
	"sync"
	"fmt"
)

func VideoClearDispatch(dc DataChan)error{
	res,err:=dbops.ReadVideoDeleteRecord(3)
	if err!=nil{
		return err
	}
	if len(res)==0{
		return errors.New("all task finish")
	}
	for _,id:=range res{
		dc<-id
	}
	return nil
}
func DeleteVideoFile(vid string)error{
	err:=os.Remove(VIDEO_DIR+vid)
	fmt.Println(VIDEO_DIR+vid)
	if err!=nil && !os.IsNotExist(err){
		return err
	}
	return nil
}
func VideoClearExecute(dc DataChan)error{
	errMap:=&sync.Map{}
	forloop:
		for{
			select {
			     case vid:=<-dc:
			     	go func(id interface{}){
						if err:=DeleteVideoFile(id.(string));err!=nil{
							errMap.Store(id,err)
							return
						}
			     		if err:=dbops.DelVideoRecord(id.(string));err!=nil{
							errMap.Store(id,err)
							return
						}
					}(vid)
			default:
				break forloop
			}
		}
		var err error
		errMap.Range(func(k,v interface{})bool{
			err=v.(error)
			if err!=nil{return false}
			return true
		})
     return err
}