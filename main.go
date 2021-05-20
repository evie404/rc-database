package main

import (
	"log"

	"github.com/rickypai/rc-database/api/http/server"
	"github.com/rickypai/rc-database/database/jsonfile"
)

func main() {
	filepath := "/tmp/db.json"
	db, err := jsonfile.NewJSONFileDatabase(filepath)
	if err != nil {
		log.Panicf("opening database: %s", err)
	}
	srv := server.NewDatabaseServer(db)

	err = srv.Listen(4000)
	if err != nil {
		log.Panicf("starting http server: %s", err)
	}
}
