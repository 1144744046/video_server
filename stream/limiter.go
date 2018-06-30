package main

import "log"

type ConnLimiter struct{
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int)*ConnLimiter{
	return &ConnLimiter{
		cc,
		make(chan int,cc),
	}
}

func (conn *ConnLimiter)GetConn()bool{
	if len(conn.bucket)>=conn.concurrentConn{
		return false
	}
	conn.bucket<-1
	return true
}

func (conn *ConnLimiter)ReleaseConn(){
      c:=<-conn.bucket
      log.Println("release conn succ",c)
}