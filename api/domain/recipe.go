package domain

import "time"

type Recipe struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	MakingTime  string    `json:"making_time,omitempty"`
	Serves      string    `json:"serves,omitempty"`
	Ingredients string    `json:"ingredients,omitempty"`
	Cost        int       `json:"cost,omitempty"`
	CreatedAt   time.Time `json:"create_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type RecipeList []Recipe
