package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rickypai/rc-database/database"
)

type DatabaseServer struct {
	database *database.Database
}

func NewDatabaseServer(db *database.Database) *DatabaseServer {
	return &DatabaseServer{
		database: db,
	}
}

func (s *DatabaseServer) Listen(port int) error {
	http.Handle("/get", &getHandler{s.database})
	http.Handle("/set", &setHandler{s.database})

	listenAddress := fmt.Sprintf(":%v", port)
	log.Printf("starting http server on %s", listenAddress)

	return http.ListenAndServe(listenAddress, nil)
}
