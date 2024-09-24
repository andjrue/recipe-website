package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserServer struct {
	address string
	userDB  *sql.DB
}

type userApiErr struct {
	Error string
}

func newUserServer(address string, db *sql.DB) *UserServer {
	return &UserServer{
		address: address,
		userDB:  db,
	}
}

func (us *UserServer) runUserServer() {
	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		err := us.handleUser(w, r) // Will write this later, same as recipe_api

		if err != nil {
			writeJson(w, http.StatusBadRequest, userApiErr{Error: err.Error()})
		}
	})

	fmt.Println("User API - Listening on Port: ", us.address) // Do we need to put this on a different port?
	http.ListenAndServe(us.address, router)
}
