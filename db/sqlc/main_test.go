package db_test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	db "simplebank/db/sqlc"
	"simplebank/util"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load configuration file or variables", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
