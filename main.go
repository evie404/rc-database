package main

import (
	"log"

	"github.com/rickypai/rc-database/database"
	"github.com/rickypai/rc-database/server"
)

func main() {
	srv := server.NewDatabaseServer(database.NewDatabase())

	err := srv.Listen(4000)
	if err != nil {
		log.Panicf("starting http server: %s", err)
	}
}
