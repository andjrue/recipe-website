import React from "react";

function DisplayRecipe({ item }) {
  return (
    <div className="recipeCard">
      <h3>item.Title</h3>
      <p>item.Ingredients</p>
      <p>item.Description</p>
      <p> item.TimeToMake</p>
    </div>
  );
}

export default DisplayRecipe;
