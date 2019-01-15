package usecase

import (
	"api/domain"
)

type RecipeRepository interface {
	Create(*domain.Recipe) error
	FindAll() (domain.RecipeList, error)
	Find(int) (domain.Recipe, error)
	// Update(int, *domain.Recipe) error
	// Delete(int) error
}
