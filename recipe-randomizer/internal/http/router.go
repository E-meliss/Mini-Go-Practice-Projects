package http

import (
	"recipe-randomizer/internal/db"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(rg *gin.RouterGroup, q *db.Queries) {
	h := &Handlers{Q: q}
	rg.GET("/recipes/random", h.Random)
	rg.GET("/recipes/:id", h.GetByID)
	rg.GET("/recipes", h.List) // optional
}
