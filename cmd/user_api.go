package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	return nil
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
	} else {
		return nil
	}

}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
