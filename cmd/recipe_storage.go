package main

import (
	"database/sql"
	"fmt"
)

type Recipe struct {
	ID           int64  `json: "id`
	Title        string `json: "title"`
	TimeToMake   string `json: "timeToMake"`
	Description  string `json: "description"`
	Ingredients  string `json: "ingredients"`
	LinkToRecipe string `json: "linkToRecipe"`
}

func newRecipe(title, timeToMake, description, ingredients, linkToRecipe string) *Recipe {
	return &Recipe{
		Title:        title,
		TimeToMake:   timeToMake,
		Description:  description,
		Ingredients:  ingredients,
		LinkToRecipe: linkToRecipe,
	}
}

/*

func randID() int {
	recipeID := rand.Intn(10000)

	return recipeID
}


I think we need to get rid of rand_id. Its needless and also creates a headache
when creating a recipe. If it's there, we'd need to query the entire DB everytime
someone adds a recipe. Feels a little silly to me to have to do that when we know the
id provided by Postgres will always be unique. May have overthought this one, but good to know.

UPDATE: This is gone. No need for it.

*/

// CREATE RECIPE TABLE

const createRecipeTable = `
CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    title TEXT,
    time_to_make TEXT,
    description TEXT,
    ingredients TEXT,
    link_to_recipe TEXT
);
`

const deleteRecipe = `
DELETE FROM recipes
WHERE id = $1
`

const getRecipe = `
SELECT * FROM recipes
WHERE id = $1
`

const insertRecipe = `
INSERT INTO
	recipes (title, time_to_make, description, ingredients, link_to_recipe)
VALUES ($1, $2, $3, $4, $5)

	`

func createRecipeTableFunc(db *sql.DB) error {
	_, err := db.Exec(createRecipeTable)

	return err
}

func insertRecipeFunc(db *sql.DB, r *Recipe) error {
	_, err := db.Exec(insertRecipe, r.Title, r.TimeToMake, r.Description, r.Ingredients, r.LinkToRecipe)
	return err
}

func getRecipeFunc(db *sql.DB, r *Recipe, id string) error { // Will pass ID in handleGetRecipe
	row := db.QueryRow(getRecipe, id)

	// Cool, this is working now w/o the rand_id's. DB is queryable by the auto generated pg IDs
	return row.Scan(&r.ID, &r.Title, &r.TimeToMake, &r.Description, &r.Ingredients, &r.LinkToRecipe)
}

func deleteRecipeFunc(db *sql.DB, id string) error {
	_, err := db.Exec(deleteRecipe, id)
	if err != nil {
		return fmt.Errorf("Issue deleting recipe: %v", err)
	}
	fmt.Printf("Recipe %v deleted!", id)
	return err
}
