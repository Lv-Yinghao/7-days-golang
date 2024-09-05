package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() {
		db.Close()
	}()

	_, _ = db.Exec("DROP TABLE IF EXISTS Users;")
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS Users (Name TEXT);")

	res, err := db.Exec("INSERT INTO USER(`Name`) values (?),(?)", "Tom", "Sam")

	if err == nil {
		affected, _ := res.RowsAffected()
		log.Println("affected:", affected)
	}

	row := db.QueryRow("SELECT * FROM Users LIMIT 1")

	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
