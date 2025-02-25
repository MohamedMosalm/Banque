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
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDeriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the DB: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
