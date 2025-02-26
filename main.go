package main

import (
	"database/sql"
	"log"

	"github.com/MohamedMosalm/banque/api"
	db "github.com/MohamedMosalm/banque/db/sqlc"
	"github.com/MohamedMosalm/banque/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	conn, err := sql.Open(config.DBDeriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
