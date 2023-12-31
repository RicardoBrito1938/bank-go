package main

import (
	"database/sql"
	"log"

	"bank.sqlc.dev/app/db/util"

	"bank.sqlc.dev/app/api"
	db "bank.sqlc.dev/app/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbSource := config.DBSource
	dbDriver := config.DBDriver
	serverAddress := config.ServerAddress

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
