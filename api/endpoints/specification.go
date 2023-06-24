package endpoints

import (
	"net/http"

	"github.com/igefined/mint/docs"
)

func (e endpoints) Specification() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Encoding", "gzip")

		_, err := w.Write(docs.Specification)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
