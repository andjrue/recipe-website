package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

type ApiError struct {
	Error string
}

func newApiServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
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
	return nil
}

// Not exactly sure how I'm going to do this
// Will probably need to add IDs to each recipe, if ID exists and POST, update the recipe
func (s *Server) handleUpdateRecipe(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeleteRecipe(w http.ResponseWriter, r *http.Request) error {
	return nil
}
