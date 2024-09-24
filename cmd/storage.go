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

func createRecipeTableFunc(db *sql.DB) error {
	_, err := db.Exec(createRecipeTable)

	return err
}

// INSERT RECIPE

const insertRecipe = `
INSERT INTO
	recipes (title, time_to_make, description, ingredients, link_to_recipe)
VALUES ($1, $2, $3, $4, $5)

	`

func insertRecipeFunc(db *sql.DB, r *Recipe) error {
	_, err := db.Exec(insertRecipe, r.Title, r.TimeToMake, r.Description, r.Ingredients, r.LinkToRecipe)
	return err
}

// GET RECIPE

const getRecipe = `
SELECT * FROM recipes
WHERE id = $1
`

func getRecipeFunc(db *sql.DB, r *Recipe, id string) error { // Will pass ID in handleGetRecipe
	row := db.QueryRow(getRecipe, id)

	// Cool, this is working now w/o the rand_id's. DB is queryable by the auto generated pg IDs
	return row.Scan(&r.ID, &r.Title, &r.TimeToMake, &r.Description, &r.Ingredients, &r.LinkToRecipe)
}

// DELETE RECIPE

const deleteRecipe = `
DELETE FROM recipes
WHERE id = $1
`

func deleteRecipeFunc(db *sql.DB, r *Recipe, id string) error {
	_, err := db.Exec(deleteRecipe, id)
	if err != nil {
		return fmt.Errorf("Issue deleting recipe: %v", err)
	}
	fmt.Print("Recipe %s deleted!", id)
	return err
}
