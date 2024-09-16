package main

import (
	"database/sql"
	"math/rand"
)

type Recipe struct {
	ID           int    `json: "id"`
	Title        string `json: "title"`
	TimeToMake   string `json: "timeToMake"`
	Description  string `json: "description"`
	Ingredients  string `json: "ingredients"`
	LinkToRecipe string `json: "linkToRecipe"`
}

func newRecipe(id int, title, description, ingredients, timeToMake, linkToRecipe string) *Recipe {
	return &Recipe{
		ID:           id,
		Title:        title,
		TimeToMake:   timeToMake,
		Description:  description,
		Ingredients:  ingredients,
		LinkToRecipe: linkToRecipe,
	}
}

func randID() int {
	recipeID := rand.Intn(10000)

	return recipeID
}

const createRecipeTable = `
CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    rand_id INT,
    title TEXT,
    time_to_make TEXT,
    description TEXT,
    ingredients TEXT,
    link_to_recipe TEXT
);
`

func createRecipeTableFunc(db *sql.DB) error {
	_, err := db.Exec(createRecipeTable)

	return err
}

const insertRecipe = `
		INSERT INTO
			recipes (rand_id, title, time_to_make, description, ingredients, link_to_recipe)
		VALUES ($1, $2, $3, $4, $5, $6)

	`

func insertRecipeFunc(db *sql.DB, r *Recipe) error {
	_, err := db.Exec(insertRecipe, r.ID, r.Title, r.TimeToMake, r.Description, r.Ingredients, r.LinkToRecipe)
	return err
}
