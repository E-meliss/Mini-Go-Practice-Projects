package http

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"recipe-randomizer/internal/db"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Q *db.Queries
}

func (h *Handlers) Random(c *gin.Context) {
	meal := strings.TrimSpace(c.Query("meal"))
	if meal == "" {
		meal = c.Query("meal_type") // alias
	}
	if meal != "" {
		meal = strings.ToLower(meal)
	}

	var maxTimePtr *int
	if mt := strings.TrimSpace(c.Query("max_time")); mt != "" {
		if n, err := strconv.Atoi(mt); err == nil && n > 0 {
			maxTimePtr = &n
		}
	}

	parseCSV := func(s string) []string {
		if strings.TrimSpace(s) == "" {
			return nil
		}
		parts := strings.Split(s, ",")
		out := make([]string, 0, len(parts))
		for _, p := range parts {
			p = strings.ToLower(strings.TrimSpace(p))
			if p != "" {
				out = append(out, p)
			}
		}
		return out
	}

	include := parseCSV(c.Query("include"))
	exclude := parseCSV(c.Query("exclude"))

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	r, err := h.Q.RandomRecipe(ctx, &meal, maxTimePtr, include, exclude)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no recipe found", "detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h *Handlers) GetByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	r, err := h.Q.GetRecipe(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h *Handlers) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	rs, err := h.Q.ListRecipes(c.Request.Context(), limit, offset)
	if err != nil {
		// TEMP: show detail to debug; remove later.
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "db error",
			"detail": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rs)
}
