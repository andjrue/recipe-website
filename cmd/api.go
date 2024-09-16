package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	db         *sql.DB
}

// post was surprisingly easy, hyst needed to add db to the server. Working.

type ApiError struct {
	Error string
}

func newApiServer(listenAddr string, db *sql.DB) *Server {
	return &Server{
		listenAddr: listenAddr,
		db:         db,
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		err := s.handleRecipe(w, r)
		if err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	})
	fmt.Println("Listening on Port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleRecipe(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		return s.handleGetRecipe(w, r)
	case "POST":
		return s.handleCreateRecipe(w, r)
	case "DELETE":
		return s.handleDeleteRecipe(w, r)
	default:
		return fmt.Errorf("Method not allowed: ", r.Method)
	}
}

func (s *Server) handleGetRecipe(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *Server) handleCreateRecipe(w http.ResponseWriter, r *http.Request) error {
	id := randID()
	recipe := newRecipe(id, "First Recipe", "", "Test", "Test", "Test")

	/*
		Todays solutions are tomorrows problems. I also realized I was doing this under getRecipe. Not ideal.

		The idea is to use random IDs for each recipe, and we'll probably have to query the DB to see if that ID already exists. If it does, we can call randID()
		again and the chances are very low that it will be another taken ID.

		Realitically, there won't be anywhere close to 10,000 recipes on the website. If that happens we have much larger problems than IDs, I'd owe Jeff so much cash.
	*/
	insertRecipeFunc(s.db, recipe)
	return writeJson(w, http.StatusOK, recipe)
}

// Not exactly sure how I'm going to do this
// Will probably need to add IDs to each recipe, if ID exists and POST, update the recipe
func (s *Server) handleUpdateRecipe(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeleteRecipe(w http.ResponseWriter, r *http.Request) error {
	return nil
}
