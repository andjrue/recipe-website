package main

import "database/sql"

type User struct {
	ID       int64  `json:"id"`
	Email string `json:"email"` // Let's make this email. I don't want to reset someones password. 
	Password string `json: "password"` // This will need to be hashed before they go into db https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60
}

// I think we can keep these this simple. I don't see a reason to add any additional fields (right now).
// Maybe we could include email or something if we wanted to update users on new recipes added but that feels
// a little out of scope. Maybe later down the road.

func newUser(em, pw string) *User {
	return &User{
		Email: em,
		Password: pw,
	}
}

const createUserTable = `
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	email varchar(50) UNIQUE NOT NULL,
	hashed_pass TEXT NOT NULL
);
`

const insertUser = `
INSERT INTO users (email, hashed_pass)
VALUES ($1, $2)
`

func createUserTableFunc(db *sql.DB) error {
	_, err := db.Exec(createUserTable)
	return err
}
