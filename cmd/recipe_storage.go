package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Recipe struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	TimeToMake   string `json:"timeToMake"`
	Description  string `json:"description"`
	Ingredients  string `json:"ingredients"`
	LinkToRecipe string `json:"linkToRecipe"`
	Email        string   `json:"email"` // We can send this from the front end.

  // Might need to store the users email in state or something. When we post, we can include it. Should be simple enough.
  // This works because the user will be signed in. For now, anyway. 
}

func newRecipe(title, timeToMake, description, ingredients, linkToRecipe, email string) *Recipe {
	return &Recipe{
		Title:        title,
		TimeToMake:   timeToMake,
		Description:  description,
		Ingredients:  ingredients,
		LinkToRecipe: linkToRecipe,
    Email: email,
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
SELECT *
FROM recipes
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
	_, err := db.Exec(insertRecipe, r.Title, r.TimeToMake, r.Description, r.Ingredients, r.LinkToRecipe, r.Email)
	return err
}

func getAllRecipesFunc(db *sql.DB) ([]Recipe, error) { // Will pass ID in handleGetRecipe
	rows, err := db.Query(getRecipe)

	if err != nil {
		log.Printf("Issue querying recipes: %v", err)
	}

	defer rows.Close()

	var res []Recipe

	for rows.Next() {
		var r Recipe
		err := rows.Scan(&r.ID, &r.Title, &r.TimeToMake, &r.Description, &r.Ingredients, &r.LinkToRecipe)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		log.Println()
		res = append(res, r)
	}
	log.Printf("Res: %v", res)
	// Cool, this is working now w/o the rand_id's. DB is queryable by the auto generated pg IDs
	return res, nil
}

func deleteRecipeFunc(db *sql.DB, id string) error {
	_, err := db.Exec(deleteRecipe, id)
	if err != nil {
		return fmt.Errorf("Issue deleting recipe: %v", err)
	}
	fmt.Printf("Recipe %v deleted!", id)
	return err
}
