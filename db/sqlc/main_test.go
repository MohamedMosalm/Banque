package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDeriver = "postgres"
	dbSource  = "postgresql://postgres:password@localhost:5432/banque?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDeriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the DB: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
