package common

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DBM *sql.DB
var errm error

//DB Connection
func connectToMySql() {
	DBM, errm = sql.Open("mysql", "root@tcp(localhost:3306)/CPP")
	if errm != nil {
		fmt.Println(errm.Error())
	}
	//defer DB.Close()

	//test connection
	errm = DBM.Ping()
	if errm != nil {
		fmt.Println(errm.Error())
	}
}
