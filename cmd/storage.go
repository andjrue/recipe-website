package main

import (
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
