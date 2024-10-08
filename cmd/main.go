/*

===============================================================
                            TO-DO's
===============================================================

1.) Get Posts connected to the DB for testing. Can do that pretty quickly. **DONE**
2.) Need to tackle update logic when that's done.
	-> I probably need to add an {id} route. That would make the update logic much more sound.
		 If we can use the ID in the route, then wouldn't need to even query the DB. We could go directly to
		 the correct recipe.

		 Lots to consider for this. It definitely won't be as easy as I'm making it sound.
3.) Handle deleting recipes. Should be easy, simple as an ID search (I think) **DONE**
4.) User creation and password hashing **DONE**
5.) Start on the front end
6.) Really figure out what we want to do here.
	-> There are about a million ways I could go with this project. I'm not entirely sure what I want to do.
	-> I also need to add usernames to the recipes and make sure each one aligns. I don't want someone else to be able to
	   delete anybodys recipe.

===============================================================
                      THINGS TO CONSIDER
===============================================================

- How do we want to do the front end?
	-> Probably react. Never used it, should be fun. From what I've seen it looks straightforward enough.
	-> I want users to be able to create folder(s). For example, I could have a folder, "Drew's Folder", that holds all of my recipes.
		 Inside of Drew's Folder, there could be folders for entrees, deserts, etc. It would be cool for each person to have a public folder,
		 that way everybody can see their recipes. It would also lead to "Hey check out what I made tonight" in the family group chat and I think
		 that would be nice.

- Do we want user creation?
 -> At present moment, I think the answer is no. There's a good chance we only have one user and can allow posts to be made from "anybody".
    In the event more people want to add in, I think it would be required to have account creation. Might need to look into JWT, etc.
    Would be nice to have it be linking a google account.
    -> Going to go with this. I think username and password will suffice. There's no need to get too crazy.
       Never going to send them an email or anything. Again, it will be very small.

- What else is needed in AWS?
	-> Definitely need to deploy this to an EC2 instance at some point. I can't have it running on my computer forever.
		-> Will most likely also require us to revisit the DB. Will need to update security group as well.

*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Printf("Error loading env: %v", envErr)
	}

	//fmt.Println("env loaded")

	db, err := sql.Open("pgx", os.Getenv("PG_DSN"))

	//fmt.Println("sql.open")
	if err != nil {
		log.Printf("Unable to connect to Database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	// fmt.Println("Pinging db")
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to recipe DB")

	err = createRecipeTableFunc(db)
	if err != nil {
		fmt.Println("Issue creating recipe table")
	} else {
		log.Println("Recipe table created or already exists")
	}

	err = createUserTableFunc(db)
	if err != nil {
		log.Println("Issue creating user table")
	}

	log.Println("User table created or already exists")

	server := newApiServer(":3000", db)
	server.Run()

}
