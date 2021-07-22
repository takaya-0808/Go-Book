package service

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBEngine *sql.DB

type SqlHandler struct {
	Conn *sql.DB
}

func DBConnect() *SqlHandler {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbProtocol := "tcp(db:3306)"
	dbName := "go_database"
	dbOption := "?parseTime=true"
	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"/"+dbName+dbOption)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
