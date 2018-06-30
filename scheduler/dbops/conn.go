package dbops

import ("database/sql"
	_ "github.com/go-sql-driver/mysql"
	"siliang/config"
)
var (
	sqlconn *sql.DB
	err error
)
func init(){

	sqlconn,err=sql.Open("mysql",config.GetDSN())
	if err!=nil{
		panic(err.Error())
	}
}
