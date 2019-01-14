package gateway

import (
	"api/domain"
	"encoding/json"
	"io"
	"net/http"
)

type RecipeList struct {
	Recipes domain.RecipeList `json:"recipes"`
}

type Error struct {
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(Error{Message: err.Error()})
	w.WriteHeader(http.StatusInternalServerError)
}

func DecodeBody(domain interface{}, body io.Reader) interface{} {
	json.NewDecoder(body).Decode(domain)
	return domain
}

func WriteRecipeList(w http.ResponseWriter, recipes domain.RecipeList) {
	json.NewEncoder(w).Encode(RecipeList{Recipes: recipes})
	w.WriteHeader(http.StatusOK)
}
