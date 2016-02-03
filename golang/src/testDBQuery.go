package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
    timestart := time.Now()

    db, err := sql.Open("mysql", "root@/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
	
	//opens 100 connections to the database
    for c := 0; c < 100; c++ {
		query, err := db.Query("SELECT idtestTable FROM testtable WHERE content = 'a test content with BMW'")
		if err != nil {
			panic(err.Error())
		}

		var col1 int
		query.Next();
		err = query.Scan(&col1);
		if err != nil {
			panic(err.Error())
		}
		defer query.Close()
    }
    timeend := time.Now()

    fmt.Print(timeend.Sub(timestart))
}