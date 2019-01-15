package controller

import (
	"api/domain"
	"api/interface/gateway"
	"api/interface/repository"
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RecipeController struct {
	Usecase usecase.RecipeUsecase
}

func NewRecipeController(sqlHandler repository.SqlHandler) *RecipeController {
	return &RecipeController{
		Usecase: usecase.RecipeUsecase{
			RecipeRepository: &repository.RecipeRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RecipeController) Create(w http.ResponseWriter, r *http.Request) {
	decoded := gateway.DecodeBody(&domain.Recipe{}, r.Body)
	recipe := decoded.(*domain.Recipe)
	if err := controller.Usecase.Create(recipe); err != nil {
		gateway.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (controller *RecipeController) ListAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := controller.Usecase.RecipeList()
	if err != nil {
		gateway.WriteError(w, err)
		return
	}
	gateway.WriteRecipeList(w, recipes)
}

func (controller *RecipeController) DetailRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	recipe, err := controller.Usecase.DetailRecipe(id)
	if err != nil {
		gateway.WriteError(w, err)
		return
	}
	gateway.WriteRecipe(w, recipe)
}
