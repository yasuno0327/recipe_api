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

func (repo *RecipeRepository) Find(id int) (domain.Recipe, error) {
	row, err := repo.Query("SELECT * FROM recipes WHERE id = ?", id)
	if err != nil {
		return domain.Recipe{}, err
	}
	defer row.Close()
	var cost int
	var title, makingTime, serves, ingredients string
	var createdAt, updatedAt []uint8
	row.Next()
	row.Scan(&id, &title, &makingTime, &serves, &ingredients, &cost, &createdAt, &updatedAt)
	recipe := domain.Recipe{
		ID:          id,
		Title:       title,
		MakingTime:  makingTime,
		Serves:      serves,
		Ingredients: ingredients,
	}
	return recipe, nil
}

// func Update() error {
// }

// func Delete() error {
// }
