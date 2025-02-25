package main

import (
	"database/sql"
	"log"

	"github.com/MohamedMosalm/banque/api"
	db "github.com/MohamedMosalm/banque/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDeriver     = "postgres"
	dbSource      = "postgresql://postgres:password@localhost:5432/banque?sslmode=disable"
	serverAddress = "0.0.0.0:9090"
)

func main() {
	var err error
	conn, err := sql.Open(dbDeriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
