package main

type User struct {
	Username string
	Password string // This will need to be hashed before they go into db
}

// I think we can keep these this simple. I don't see a reason to add any additional fields (right now).
// Maybe we could include email or something if we wanted to update users on new recipes added but that feels
// a little out of scope. Maybe later down the road.
