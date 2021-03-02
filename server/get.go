package server

import (
	"fmt"
	"net/http"

	"github.com/rickypai/rc-database/database"
)

type getHandler struct {
	database *database.Database
}

func (h *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		// handle
	}

	value, err := h.database.Get(key)
	if err != nil {
		// handle
	}

	if value == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, "%s", value)
}
