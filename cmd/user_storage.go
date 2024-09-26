package main

import "database/sql"

type User struct {
	Username string `json:"username"`
	Password string `json: "password"` // This will need to be hashed before they go into db https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60
}

// I think we can keep these this simple. I don't see a reason to add any additional fields (right now).
// Maybe we could include email or something if we wanted to update users on new recipes added but that feels
// a little out of scope. Maybe later down the road.

func newUser(un, pw string) *User {
	return &User{
		Username: un,
		Password: pw,
	}
}

const createUserTable = `
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	username varchar(50) UNIQUE NOT NULL,
	hashed_pass TEXT NOT NULL
);
`

const insertUser = `
INSERT INTO users (username, hashed_pass)
VALUES ($1, $2)
`

func createUserTableFunc(db *sql.DB) error {
	_, err := db.Exec(createUserTable)
	return err
}
