/* 
 *  What we need to do:
 *
 *  The front-end is mostly working now and I'm beginning to understand React. 
 *  Because of that, I think we should turn our eyes back to the backend. 
 *
 *  1.) Get user auth working. I think JWT will be the best way to do this, but I honestly have no idea. 
 *   -> Need to do more research on this.
 *   -> I need to connect a user to the recipe uploads anyway, may as well do it now.
 *      -> 10/7 - I've changed my mind about this. I need to build the login and post recipes functionality on the front end before I can do this. 
 *                That way I can actually test the functionality when I'm signed in to the website. Otherwise it would be a guessing game if it's actually working.
 *
 *
 *                So what we will do: 
 *
 *                Get user sign in working. After that, we tackle deletion issues. We also only want users to be able to update their own recipes. 
 *                Still not 100% sure on JWT, but we can add that later if we want. For now, I think it will be as simple as checking if the username of the person
 *                that uploaded the recipe matches the person with the username trying to delete/update the recipe. 
 *
 *  2.) Once that's done, we can go back to the front end and finish it out. Step 1 will be to get user sign-in to work and actually 
 *      write to the DB. Passwords are already hashed, so we really just need
 *
 *  I'm not exactly how the front end is going to end up. I've already had to make some heavy adjuments to the backend and we just started. 
 *  I expected that to happen, but learning is still painful. Either way, I think there will be additions and deletions to both the front and back end 
 *  as I go. That's to be expected, and it shouldn't be nearly as painful as it was before. 
 *
 *  In short, I'll have to figure this part out later. More important things right now. I also suck as CSS and think it's the worst "language" out there.
 *  Either way, going to need to figure it out. That will probably be the last thing that gets done. Hopefully bootstrap has some nice templates we can
 *  use. 
 * */






// import DisplayRecipe from "./RecipeCards";
import { useState } from 'react';
import CreateAccountButton from "./AccountButton";
import LoginForm from "./CreateAccount";


import 'bootstrap/dist/css/bootstrap.min.css';


function App() {

    const [showCreateAccount, setShowCreateAccount] = useState(false);

    return (
        <div className="App">
            <CreateAccountButton onClick = {() => setShowCreateAccount(true)} />
            {showCreateAccount && <LoginForm />}
        </div>
    );
}

export default App;
