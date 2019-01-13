package repository

import (
	"api/domain"
	"log"
)

type RecipeRepository struct {
	SqlHandler
}

func (repo *RecipeRepository) Create(recipe *domain.Recipe) error {
	query, err := repo.Execute(
		"INSERT INTO recipes (title, making_time, serves, ingredients, cost) VALUES (?, ?, ?, ?, ?)",
		recipe.Title,
		recipe.MakingTime,
		recipe.Serves,
		recipe.Ingredients,
		recipe.Cost,
	)
	if err != nil {
		return err
	}
	log.Println(query)
	return nil
}

// func FindAll() (domain.RecipeList, error) {
// }

// func Find() (domain.Recipe, error) {
// }

// func Update() error {
// }

// func Delete() error {
// }
