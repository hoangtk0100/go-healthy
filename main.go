package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/hoangtk0100/go-healthy/api"
	db "github.com/hoangtk0100/go-healthy/db/sqlc"
	"github.com/hoangtk0100/go-healthy/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
