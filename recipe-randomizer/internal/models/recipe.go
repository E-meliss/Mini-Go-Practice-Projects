package models

import "time"

type Recipe struct {
	ID              int64     `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Description     *string   `db:"description" json:"description,omitempty"`
	MealType        string    `db:"meal_type" json:"mealType"`
	CookTimeMinutes int       `db:"cook_time_minutes" json:"cookTimeMinutes"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
}

type Ingredient struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type RecipeIngredient struct {
	RecipeID     int64  `db:"recipe_id" json:"recipeId"`
	IngredientID int64  `db:"ingredient_id" json:"ingredientId"`
	Quantity     string `db:"quantity" json:"quantity"`
}

type IngredientLine struct {
	Name     string `db:"name" json:"name"`
	Quantity string `db:"quantity" json:"quantity"`
}

type RecipeWithIngredients struct {
	Recipe      Recipe           `json:"recipe"`
	Ingredients []IngredientLine `json:"ingredients"`
}
