package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	db "recipe-randomizer/internal/db"
	httpx "recipe-randomizer/internal/http"
	"recipe-randomizer/internal/web"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "file:./recipes.db?_fk=1&_busy_timeout=5000"
	}

	sqlDB, err := db.Open(dsn)
	if err != nil {
		log.Fatal("open db:", err)
	}
	defer sqlDB.Close()

	if err := db.AutoMigrateAndSeed(sqlDB); err != nil {
		log.Fatal("migrate/seed:", err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	q := db.NewQueries(sqlDB)

	// API routes
	api := r.Group("/api")
	httpx.RegisterAPI(api, q)

	// Minimal HTML page using Go templates
	r.GET("/", func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		web.RenderRecipePage(c.Writer, c.Request)
	})

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Println("listening on :" + port)
	log.Fatal(s.ListenAndServe())
}
