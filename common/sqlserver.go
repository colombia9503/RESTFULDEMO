package common

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB
var err error

//DB Connection
func connectToSqlServer() {
	connString := "sqlserver://sa:accesobd@192.168.0.102:1433?database=prueba"
	DB, err = sql.Open("mssql", connString)
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer DB.Close()

	//test connection
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
}
