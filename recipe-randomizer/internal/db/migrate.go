package db

import (
	"context"
	_ "embed"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var schemaSQL string

//go:embed seed.sql
var seedSQL string

func AutoMigrateAndSeed(db *sqlx.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := db.ExecContext(ctx, schemaSQL); err != nil {
		return err
	}

	// seed only if empty
	var n int
	if err := db.GetContext(ctx, &n, `SELECT COUNT(1) FROM recipes`); err == nil && n == 0 {
		stmts := strings.Split(seedSQL, ";")
		for _, s := range stmts {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if _, err := db.ExecContext(ctx, s); err != nil {
				return err
			}
		}
	}
	return nil
}
