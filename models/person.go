package models

import _ "github.com/mattn/go-sqlite3"

type Person struct {
	Id         int
	First_name string
	Last_name  string
	Email      string
	Ip_address string
}
