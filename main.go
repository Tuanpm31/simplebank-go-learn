package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}

	DB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}
	store := db.NewStore(DB)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start  server", err)
	}
}
