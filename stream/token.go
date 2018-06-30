package main

import (
	"time"
	"math/rand"
	"strconv"
)

var tokenMap map[string]int
func init(){
	tokenMap=make(map[string]int)
}

func getToken()string{
	t:=time.Now()
	timestamp:=strconv.FormatInt(t.UTC().UnixNano(), 10)
	tt:=timestamp+strconv.Itoa(rand.Int())
	tokenMap[tt]=1
	return tt
}
func checkToken(key string)bool{
	if _, ok := tokenMap[key]; ok {
		//存在
		return true
	}
	return false
}
func deleteToken(key string){
	delete(tokenMap,key)
}