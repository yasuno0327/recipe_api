package usecase

import (
	"api/domain"
)

type RecipeUsecase struct {
	RecipeRepository RecipeRepository
}

func (usecase *RecipeUsecase) Create(recipe *domain.Recipe) error {
	err := usecase.RecipeRepository.Create(recipe)
	return err
}

func (usecase *RecipeUsecase) RecipeList() (domain.RecipeList, error) {
	recipes, err := usecase.RecipeRepository.FindAll()
	return recipes, err
}

// func (usecase *RecipeUsecase) Find(id int) (*domain.Recipe, error) {
// 	recipe, err := usecase.RecipeRepository.Find(id)
// 	return recipe, err
// }

// func (usecase *RecipeUsecase) Update(id int, recipe *domain.Recipe) error {
// 	err := usecase.RecipeRepository.Update(id, recipe)
// 	return err
// }

// func (usecase *RecipeUsecase) Destroy(id int) error {
// 	err := usecase.RecipeRepository.Delete(id)
// 	return err
// }
