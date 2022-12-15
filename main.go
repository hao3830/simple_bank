package main

import (
	"database/sql"
	"log"

	"github.com/hao3830/simple_bank/api"
	db "github.com/hao3830/simple_bank/db/sqlc"
	"github.com/hao3830/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".", "app", "env")

	if err != nil {
		log.Fatal("Cannot load config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
