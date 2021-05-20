package server

import (
	"fmt"
	"log"
	"net/http"
)

type DataReader interface {
	Get(key string) ([]byte, error)
}

type DataWriter interface {
	Set(key string, value []byte) error
}

type DataReaderWriter interface {
	DataReader
	DataWriter
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
