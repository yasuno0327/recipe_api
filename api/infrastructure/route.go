package infrastructure

import (
	"api/interface/controller"
	"api/interface/repository"
	"net/http"

	"github.com/gorilla/mux"
)

var sqlHandler repository.SqlHandler

func Run() {
	r := mux.NewRouter()
	sqlHandler = NewSqlHandler()
	recipes(r)
	http.ListenAndServe(":3000", r)
}

func recipes(r *mux.Router) {
	recipePrefix := "/recipes"
	recipeController := controller.NewRecipeController(sqlHandler)
	r.HandleFunc(recipePrefix, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			recipeController.ListAll(w, r)
		case http.MethodPost:
			recipeController.Create(w, r)
		default:
			w.WriteHeader(404)
		}
	})

	r.HandleFunc(recipePrefix+"/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		case http.MethodPatch:
		case http.MethodDelete:
		default:
			w.WriteHeader(404)
		}
	})
}
