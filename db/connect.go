package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root:tri123@tcp(localhost:3305)/thedove")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
