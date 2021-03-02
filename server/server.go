package server

import (
	"fmt"
	"log"
	"net/http"
)

type DataReaderWriter interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
type DatabaseServer struct {
	mux *http.ServeMux
}

func NewDatabaseServer(db DataReaderWriter) *DatabaseServer {
	mux := http.NewServeMux()
	mux.Handle("/get", &getHandler{db})
	mux.Handle("/set", &setHandler{db})

	return &DatabaseServer{
		mux: mux,
	}
}

func (s *DatabaseServer) Listen(port int) error {
	listenAddress := fmt.Sprintf(":%v", port)
	log.Printf("starting http server on %s", listenAddress)

	return http.ListenAndServe(listenAddress, s.mux)
}
