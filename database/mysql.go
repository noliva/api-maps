package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:root@tcp(192.168.99.100:3306)/api_maps")
	if err != nil {
		log.Fatalln("Error initializing DB: ", err)
	}
}

func Connector() *sqlx.DB {
	return db
}