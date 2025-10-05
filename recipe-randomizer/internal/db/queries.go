package db

import (
	"context"
	"strings"
	"time"

	"recipe-randomizer/internal/models"

	"github.com/jmoiron/sqlx"
)

type Queries struct {
	db *sqlx.DB
}

func NewQueries(db *sqlx.DB) *Queries { return &Queries{db: db} }

func (q *Queries) GetRecipe(ctx context.Context, id int64) (models.RecipeWithIngredients, error) {
	var r models.Recipe
	if err := q.db.GetContext(ctx, &r, `
		SELECT id,title,description,meal_type,cook_time_minutes,created_at
		FROM recipes WHERE id = ?`, id); err != nil {
		return models.RecipeWithIngredients{}, err
	}
	ings := []models.IngredientLine{}
	if err := q.db.SelectContext(ctx, &ings, `
		SELECT i.name, COALESCE(ri.quantity,'') AS quantity
		FROM recipe_ingredients ri
		JOIN ingredients i ON i.id = ri.ingredient_id
		WHERE ri.recipe_id = ?
		ORDER BY i.name ASC`, id); err != nil {
		return models.RecipeWithIngredients{}, err
	}
	return models.RecipeWithIngredients{Recipe: r, Ingredients: ings}, nil
}

func (q *Queries) RandomRecipe(ctx context.Context, meal *string, maxTime *int, include, exclude []string) (models.RecipeWithIngredients, error) {
	var (
		sql = `
SELECT r.id, r.title, r.description, r.meal_type, r.cook_time_minutes, r.created_at
FROM recipes r
WHERE 1=1`
		args []any
	)

	if meal != nil && *meal != "" {
		sql += ` AND r.meal_type = ?`
		args = append(args, strings.ToLower(*meal))
	}
	if maxTime != nil && *maxTime > 0 {
		sql += ` AND r.cook_time_minutes <= ?`
		args = append(args, *maxTime)
	}

	if len(include) > 0 {
		for i := range include {
			include[i] = strings.ToLower(strings.TrimSpace(include[i]))
		}
		sql += `
AND r.id IN (
  SELECT ri.recipe_id
  FROM recipe_ingredients ri
  JOIN ingredients i ON i.id = ri.ingredient_id
  WHERE i.name IN (` + placeholders(len(include)) + `)
  GROUP BY ri.recipe_id
  HAVING COUNT(DISTINCT i.name) = ?
)`
		for _, v := range include {
			args = append(args, v)
		}
		args = append(args, len(include))
	}

	if len(exclude) > 0 {
		for i := range exclude {
			exclude[i] = strings.ToLower(strings.TrimSpace(exclude[i]))
		}
		sql += `
AND r.id NOT IN (
  SELECT ri.recipe_id
  FROM recipe_ingredients ri
  JOIN ingredients i ON i.id = ri.ingredient_id
  WHERE i.name IN (` + placeholders(len(exclude)) + `)
)`
		for _, v := range exclude {
			args = append(args, v)
		}
	}

	sql += ` ORDER BY RANDOM() LIMIT 1`

	var r models.Recipe
	if err := q.db.GetContext(ctx, &r, sql, args...); err != nil {
		return models.RecipeWithIngredients{}, err
	}
	return q.GetRecipe(ctx, r.ID)
}

func placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.TrimRight(strings.Repeat("?,", n), ",")
}

func (q *Queries) ListRecipes(ctx context.Context, limit, offset int) ([]models.Recipe, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	rs := []models.Recipe{}
	err := q.db.SelectContext(ctx, &rs, `
		SELECT id,title,description,meal_type,cook_time_minutes,created_at
		FROM recipes
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?`, limit, offset)
	return rs, err
}

func withTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}
