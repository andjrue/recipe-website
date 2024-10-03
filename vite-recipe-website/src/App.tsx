import { useState, useEffect } from "react";

interface Recipe {
  // This needs to match the backend res
  ID: number;
  Title: string;
  TimeToMake: string;
  Description: string;
  Ingredients: string;
  LinkToRecipe: string;
}

const BASE_URL = "http://localhost:3000";

function App() {
  const [recipes, setRecipes] = useState<Recipe[]>([]);

  useEffect(() => {
    const fetchRecipes = async () => {
      const response = await fetch(`${BASE_URL}/recipes`);
      const recipes = (await response.json()) as Recipe[];
      setRecipes(recipes);
    };

    fetchRecipes();
  }, []);

  return (
    <div className="Test">
      <h1>Fetching Recipes</h1>
      <ul>
        {recipes.map((recipe) => {
          return <li key={recipe.ID}>{recipe.Title}</li>;
        })}
      </ul>
    </div>
  );
}

export default App;
