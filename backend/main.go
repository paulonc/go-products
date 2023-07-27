package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/paulonc/go-products/backend/api"
	db "github.com/paulonc/go-products/backend/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_products?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	store := db.ExecuteNewStore(conn)
	server := api.InstanceServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("api started with error:", err)
	}
}