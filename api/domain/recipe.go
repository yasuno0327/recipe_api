package domain

import "time"

type Recipe struct {
	ID          string
	Title       string
	MakingTime  string
	Serves      string
	Ingredients string
	Cost        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
