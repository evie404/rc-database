package server

import (
	"fmt"
	"net/http"
)

type getHandler struct {
	database DataReaderWriter
}

func (h *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte("error: key parameter not set"))

		return
	}

	value, err := h.database.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(fmt.Sprintf("error: error getting key `%s`: %s", key, err)))

		return
	}

	if value == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, "%s", value)
}
