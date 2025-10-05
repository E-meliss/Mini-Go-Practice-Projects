# Mini Go Practice Projects

Recipe Randomizer

A Go web application that serves random recipes from a SQLite database.
You can filter results by meal type, maximum cooking time, and included or excluded ingredients.
The project also includes a minimal HTML interface built using Go templates for quick testing.

Overview

This project demonstrates how to build a modular Go application with:

Filtering queries using dynamic SQL

Table joins (recipes ↔ ingredients)

Randomized result selection

RESTful JSON APIs with query parameters

Templated HTML rendering with html/template

SQLite database setup, migration, and seeding

Project Structure
recipe-randomizer/
├─ cmd/
│  └─ api/
│     └─ main.go                # Application entry point
├─ internal/
│  ├─ db/                       # Database setup, queries, schema, and seeding
│  ├─ http/                     # HTTP routes and handlers (Gin)
│  ├─ models/                   # Data models
│  └─ web/                      # HTML templates and rendering
├─ go.mod
├─ go.sum
├─ .env
└─ README.md

Setup
1. Clone the repository
git clone https://github.com/yourname/recipe-randomizer.git
cd recipe-randomizer

2. Initialize dependencies
go mod tidy

3. Create a .env file
PORT=8080
DATABASE_URL=file:./recipes.db?_fk=1&_busy_timeout=5000

4. Run the application
go run ./cmd/api


The server will start at:

http://localhost:8080

API Endpoints
GET /api/recipes/random

Returns a random recipe matching the given filters.

Query parameters:

Parameter	Type	Description
meal	string	Filter by meal type (breakfast, lunch, dinner, etc.)
max_time	int	Maximum cooking time in minutes
include	string	Comma-separated ingredients that must be included
exclude	string	Comma-separated ingredients that must be excluded

Example

/api/recipes/random?meal=dinner&include=tomato,basil&exclude=garlic


Response

{
  "recipe": {
    "id": 3,
    "title": "Pasta Pomodoro",
    "mealType": "dinner",
    "cookTimeMinutes": 20
  },
  "ingredients": [
    { "name": "tomato", "quantity": "5" },
    { "name": "basil", "quantity": "handful" },
    { "name": "pasta", "quantity": "150 g" }
  ]
}
