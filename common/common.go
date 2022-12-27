package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Database() {

	var err error
	db, err := sql.Open("postgres", "user=postgres  password=eyuvaeyuva1515  dbname=postgres  sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("db connected")
		Db = db
	}

}
