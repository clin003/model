package main

import (
	"fmt"

	"gitee.com/lyhuilin/model"
)

func main() {
	dbtype, dbfile := "", ""
	dbuser, dbpass := "", ""
	dbhost, dbport, dbname := "", "", ""
	err := model.Init(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	fmt.Println(err)
	// var db *model.Database
	db, err := model.New(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	fmt.Println(db, err)
}
