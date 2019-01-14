package repository

import (
	"api/domain"
)

type RecipeRepository struct {
	SqlHandler
}

func (repo *RecipeRepository) Create(recipe *domain.Recipe) error {
	_, err := repo.Execute(
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
	return nil
}

func (repo *RecipeRepository) FindAll() (recipes domain.RecipeList, err error) {
	rows, err := repo.Query("SELECT * FROM recipes")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, cost int
		var title, makingTime, serves, ingredients string
		var createdAt, updatedAt []uint8
		if err := rows.Scan(&id, &title, &makingTime, &serves, &ingredients, &cost, &createdAt, &updatedAt); err != nil {
			panic(err)
			continue
		}
		recipe := domain.Recipe{
			ID:          id,
			Title:       title,
			MakingTime:  makingTime,
			Serves:      serves,
			Ingredients: ingredients,
		}
		recipes = append(recipes, recipe)
	}
	return
}

// func Find() (domain.Recipe, error) {
// }

// func Update() error {
// }

// func Delete() error {
// }
