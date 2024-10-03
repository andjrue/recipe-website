package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		return s.handleGetUser(w, r)
	case "POST":
		return s.handleCreateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	default:
		return fmt.Errorf("Users - Method not allowed: ", r.Method)
	}
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) error {

	log.Printf("Request: %v", r)

	db := s.db
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("entered ID: ", id)
	intId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf("Id still a strang: %v", intId)
	}

	log.Printf("ID successfully converted to int")

	u := &User{} // Need to store the result here

	if err := getUserFunc(db, u, id); err != nil {
		return fmt.Errorf("Error fetching user: %v\n", err)
	}

	return writeJson(w, http.StatusOK, u)
}

const getUserQuery = `
SELECT id, username FROM users
WHERE id = $1

`

func getUserFunc(db *sql.DB, u *User, id string) error {
	row := db.QueryRow(getUserQuery, id)
	return row.Scan(&u.ID, &u.Username)
}

func getAllUsersFunc(db *sql.DB, u *User) error {
	const getAllUsers = `
	SELECT username
	FROM users;
	`

	row := db.QueryRow(getAllUsers)
	return row.Scan(&u.Username)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) error {

	const queryUserDB = `
	SELECT COUNT(*) FROM users
	WHERE username = $1
	`

	// This will eventually be a form

	db := s.db
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("Issue loading user env", envErr)
	}
	user := newUser("drew_hash_and_pass_test_6", os.Getenv("TEST_PASS"))

	usernameExistsCh := make(chan bool)
	passwordLengthCheck := make(chan bool)

	// Doing these checks concurrently is a little silly, but I've been reading concurrency in go and wanted to do it.
	// This is more about having fun and learning than building the most efficient application of all time.

	go func() {
		start := time.Now()
		log.Println("Username check for: ", user.Username)
		var count int // We can count rows here. If > 0, we know it exists
		err := db.QueryRow(queryUserDB, user.Username).Scan(&count)
		if err != nil {
			log.Println("Email check is hella broke: ", err)
		}
		usernameExistsCh <- count > 0
		log.Println("User check complete")
		elapsed := time.Since(start)
		log.Println("User check running for: ", elapsed)
	}()

	go func() {
		start := time.Now()
		log.Println("Pass check for: ", user.Username)
		lenPass := len(user.Password)

		if lenPass < 8 {
			// This will throw some kind of error for the user
			// fmt.Println("Stuck here 2")
			passwordLengthCheck <- false
			elapsed := time.Since(start)
			log.Println("Failed pass check ran for: ", elapsed)
			http.Error(w, "Password is too short. It needs to be 8 characters or more.", http.StatusUnauthorized)

			return
		}
		// fmt.Println("Stuck here 1")
		elapsed := time.Since(start)
		log.Println("Successful pass check ran for: ", elapsed)
		log.Println("Password accepted for User: ", user.Username)
		passwordLengthCheck <- true
	}()

	usernameExists := <-usernameExistsCh
	succPass := <-passwordLengthCheck

	close(usernameExistsCh)
	close(passwordLengthCheck)

	if usernameExists {
		log.Println("Username already exists successfully captured: \n", user.Username)
		http.Error(w, "Username already exists", http.StatusConflict)
		return nil
	}

	if succPass {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password: %v\n", err)
			return err
		}
		_, err = db.Exec(insertUser, user.Username, hashedPassword)
		user.Password = string(hashedPassword) // Test, don't really want passwords being sent through API calls
		log.Println("User created: \n", user)
	}

	if !usernameExists && succPass {
		return writeJson(w, http.StatusOK, user)
	}

	return nil // We dont really need to return anything here. If we havent already returned an error than were good.

}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	// If user name == user name, logic
	return nil
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil

	/*
		A few ideas here:

		1.) We will 100% want to make sure that the username of the user requesting the deletion matches up with the username of the
		persone being deleted.
		2.) I also think it would be good practice for them to enter their password before nuking their account

	*/
}
