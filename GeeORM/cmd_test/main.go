package main

import (
	"fmt"
	GeeORM "geeorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := GeeORM.NewEngine("sqlite3", "gee.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXIST User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	res, _ := s.Raw("INSERT INTO User(`Name`) values (?),(?)", "Tom", "Peter").Exec()

	cnt, _ := res.RowsAffected()
	fmt.Println(cnt)
}
