package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// var dB *sql.DB
type dBHandler struct {
	dB *sql.DB
}

func main() {
	// Capture connection properties
	cnfig := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	//Get a database handle
	dB, err := sql.Open("mysql", cnfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	ErrTryPing := dB.Ping()
	if ErrTryPing != nil {
		log.Fatal(ErrTryPing)
	}
	fmt.Println("Connection to Database. Success")

}
