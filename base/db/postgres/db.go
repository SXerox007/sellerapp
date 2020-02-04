package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgespswd"
	dbname   = "sellerapp"
)

func DBConnecting() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetClient() *sql.DB {
	return db
}
