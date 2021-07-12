package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user     = "root"
	db_password = "tri123"
	db_address  = "127.0.0.1"
	db_name     = "thedove"
)

var Db *sql.DB

func Connect() {
	s := fmt.Sprintf("%s:%s@tcp(%s:3305)/%s", db_user, db_password, db_address, db_name)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
