package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB
var dataBase = "root:mysqlpassword@tcp(127.0.0.1:3306)/?loc=Local&parseTime=true"

func mysqlInit() {
	var err error
	DB, err = sql.Open("mysql", dataBase)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln("ping db fail:", err)
	}
}

func main43() {
	mysqlInit()

	stat, _ := DB.Prepare("select 1")

	for {
		execSql(stat)
		time.Sleep(60 * time.Second)
	}
}

func execSql(stat *sql.Stmt) {
	_, err := stat.Exec()
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	log.Println("success")
}
