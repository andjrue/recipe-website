package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		fmt.Println("Issue loading user env", envErr)
	}
	user := newUser("drew_test12", os.Getenv("TEST_PASS"))

	usernameExistsCh := make(chan bool)
	// passwordLengthCheck := make(chan bool)

	go func() {
		var count int // We can count rows here. If > 0, we know it exists
		err := db.QueryRow(queryUserDB, user.Username).Scan(&count)
		if err != nil {
			log.Println("Email check is hella broke: ", err)
		}
		usernameExistsCh <- count > 0
	}()

	usernameExists := <-usernameExistsCh

	close(usernameExistsCh)

	if usernameExists {
		fmt.Println("Username already exists successfully captured: %v", user.Username)
		http.Error(w, "Username already exists", http.StatusConflict)
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		return err
	}

	_, err = db.Exec(insertUser, user.Username, hashedPassword)
	fmt.Printf("User created: %v\n", user)

	return writeJson(w, http.StatusOK, user)
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
