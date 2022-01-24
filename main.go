package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	fmt.Println("Hellow World")

	condb, errdb := sql.Open("mssql", "server=10.253.12.19;user id=sa;password=tlstprP1@#;")
	if errdb != nil {
		fmt.Println("  Error open db:", errdb.Error())
	}

	defer condb.Close()
}
