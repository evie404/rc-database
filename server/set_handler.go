package server

import (
	"fmt"
	"net/http"
)

type setHandler struct {
	database DataReaderWriter
}

func (h *setHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) != 1 {
		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(fmt.Sprintf("error: expect only one key in query params. got %v", len(r.URL.Query()))))

		return
	}

	for key, values := range r.URL.Query() {
		if len(values) != 1 {
			w.WriteHeader(http.StatusBadRequest)

			w.Write([]byte(fmt.Sprintf("error: expect only one value on key `%s`. got %v", key, len(values))))

			return
		}

		err := h.database.Set(key, []byte(values[0]))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			w.Write([]byte(fmt.Sprintf("error: error writing key `%s`: %s", key, err)))

			return
		}
	}
}
