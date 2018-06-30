package config

import (
	"os"
	"github.com/gin-gonic/gin/json"
)

type Configuration struct {
	LBAddr string `json:"lbaddr"`
	DBHost string `json:"db_host"`
	DBUser string `json:"db_user"`
	DBPwd string `json:"db_pwd"`
	DBPort string `json:"db_port"`
	DBDatabase string `json:"db_database"`
}
var configuration *Configuration
func init(){
	file,_:=os.Open("./conf.json")
	defer file.Close()
	decode:=json.NewDecoder(file)
	configuration=&Configuration{}
	err:=decode.Decode(configuration)
	if err!=nil{
		panic(err)
	}
}

func GetLBAddr()string{
	return configuration.LBAddr
}
func GetDBHost()string{
	return configuration.DBHost
}
func GetDBUser()string{
	return configuration.DBUser
}
func GetDBPwd()string{
	return configuration.DBPwd
}
func GetDBPort()string{
	return configuration.DBPort
}
func GetDBDatabase()string{
	return configuration.DBDatabase
}
func GetDSN() string {
	dsn:=GetDBUser()+":"+GetDBPwd()+ "@tcp("+GetDBHost()+":"+
		GetDBPort()+")/"+GetDBDatabase()+"?charset=utf8"
		return dsn
}