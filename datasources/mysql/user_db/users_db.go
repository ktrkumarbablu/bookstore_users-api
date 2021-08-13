package userdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ClientDB *sql.DB
)

func init() {
	dataSources := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "password", "127.0.0.1:3306", "user_db")
	var err error
	ClientDB, err = sql.Open("mysql", dataSources)
	if err != nil {
		panic(err)
	}
	if err = ClientDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully connected")

}
