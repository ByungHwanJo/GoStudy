package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func main() {
	fmt.Println("Hellow World")

	err := ReadTest("server=10.253.12.19;user id=sa;password=tlstprP1@#;")
	if err != nil {
		log.Fatal("Error creating Employee: ", err.Error())
	}
}

func ReadTest(conn string) error {
	db, errdb := sql.Open("mssql", conn)
	if errdb != nil {
		fmt.Println("  Error open db:", errdb.Error())
	}

	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := fmt.Sprintf("SELECT CARD_NO_FROM, CARD_NO_TO FROM CloudPOS_Common.dbo.SPNT_CARD_BIN8;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var CARD_NO_FROM, CARD_NO_TO string

		// Get values from row.
		err := rows.Scan(&CARD_NO_FROM, &CARD_NO_TO)
		if err != nil {
			return err
		}

		fmt.Printf("CARD_NO_FROM: %s, CARD_NO_TO: %s\n", CARD_NO_FROM, CARD_NO_TO)
		count++
	}
	defer db.Close()
	return nil
}
