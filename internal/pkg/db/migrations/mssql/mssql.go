package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var database = "social-media-db"

func ConnectoMSSQL() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;Trusted_Connection=True;",
		server, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
}
