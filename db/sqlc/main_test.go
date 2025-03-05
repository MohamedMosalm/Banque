package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/MohamedMosalm/banque/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	configPaths := []string{
		"../..",
		"../../.github/workflows",
	}

	var config util.Config
	var err error

	for _, path := range configPaths {
		config, err = util.LoadConfig(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	testDB, err = sql.Open(config.DBDeriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
