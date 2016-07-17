package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/turnkey-commerce/go-security-examples/database"
	"github.com/turnkey-commerce/go-security-examples/utility"
)

func main() {
	// Setup the main db.
	_, err := database.InitializeDB("database/sql_injection.db")
	utility.Check(err)

}
