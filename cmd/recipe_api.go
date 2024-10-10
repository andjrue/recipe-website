package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	listenAddr string
	db         *sql.DB
}

// post was surprisingly easy, just needed to add db to the server. Working.

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

	router.HandleFunc("/recipes/{id}", func(w http.ResponseWriter, r *http.Request) {
		err := s.handleRecipe(w, r)
		if err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	})

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		err := s.handleUser(w, r)
		if err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	})

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		err := s.handleUser(w, r)
		if err != nil {
			log.Printf("AAAAAAAAAAAAA")
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Recipe API - Listening on Port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, handler)
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
		return fmt.Errorf("Recipes - Method not allowed: ", r.Method)
	}
}

func (s *Server) handleGetRecipe(w http.ResponseWriter, r *http.Request) error {
	// vars := mux.Vars(r)
	// id := vars["id"]
	// I might add ID back to this, but for now let's just worry about getting all of them.

	// rec := &Recipe{} // Need to store the result here

	res, err := getAllRecipesFunc(s.db)

	if err != nil {
		log.Fatal(err)
	}

	return writeJson(w, http.StatusOK, res)
}

func (s *Server) handleCreateRecipe(w http.ResponseWriter, r *http.Request) error {
	// Eventually this will be user inputted information. Most likely through a form.
	recipe := newRecipe("First Recipe", "", "Test", "Test", "Test", "Test Email")
	insertRecipeFunc(s.db, recipe)
	return writeJson(w, http.StatusOK, recipe)
}

func (s *Server) handleUpdateRecipe(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeleteRecipe(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	return deleteRecipeFunc(s.db, id)
}
