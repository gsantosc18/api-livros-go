package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string = "root:root@tcp(localhost:3306)/aulago"

var instance *sql.DB

func GetInstance() *sql.DB {
	if instance == nil {
		db, err := sql.Open("mysql", dsn)

		if err != nil {
			log.Fatal("erro na comunicação com o banco. ", err.Error())
		}

		if err := db.Ping(); err != nil {
			log.Fatal("falha na coneção com o banco. ", err.Error())
		}

		instance = db
	}

	return instance
}
