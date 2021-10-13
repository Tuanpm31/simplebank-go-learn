package db_test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	db "simplebank/db/sqlc"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
