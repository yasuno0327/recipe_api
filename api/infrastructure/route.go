package infrastructure

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	recipes(r)
	http.ListenAndServe(":3000", nil)
}

func recipes(r *mux.Router) {
	recipePrefix := "/recipes"
	r.HandleFunc(recipePrefix, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		case http.MethodPost:
		default:
			w.WriteHeader(404)
		}
	})

	r.HandleFunc(recipePrefix+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		case http.MethodPatch:
		case http.MethodDelete:
		default:
			w.WriteHeader(404)
		}
	})
}
