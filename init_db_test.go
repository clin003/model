package model

import (
	"testing"
)

func TestInit(t *testing.T) {
	dbtype, dbfile := "", ""
	dbuser, dbpass := "", ""
	dbhost, dbport, dbname := "", "", "test"
	err := Init(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	t.Log(err)

	dbtype, dbfile = "", ""
	dbuser, dbpass = "", ""
	dbhost, dbport, dbname = "", "", "test"
	sqlite, err := New(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	t.Log(sqlite, err)

	dbtype, dbfile = "mysql", ""
	dbuser, dbpass = "root", ""
	dbhost, dbport, dbname = "127.0.0.1", "3306", "test"
	mysql, err := New(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	t.Log(mysql, err)

	dbtype, dbfile = "postgres", ""
	dbuser, dbpass = "postgres", ""
	dbhost, dbport, dbname = "127.0.0.1", "5432", "test"
	postgres, err := New(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	t.Log(postgres, err)
}
