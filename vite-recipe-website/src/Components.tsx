import React from "react";
import { useState, useEffect } from "react";
import Card from 'react-bootstrap/Card';


function DisplayRecipe() {
  const [recipe, setRecipes] = useState([])

  useEffect(() => {
        fetch("http://localhost:3000/recipes")
            .then(response => response.json())
            .then(data => setRecipes(data));
    }, []); 

  return (
    <div>
      {recipe.map((r, index) => (
        <Card key = {index} style={{ width: '18rem', marginBottom: '1rem' }}>
          <Card.Body>
            <Card.Title>{r.Title}</Card.Title>
            <Card.Text>{r.Description}</Card.Text>
          </Card.Body>
        </Card>
      ))}
    </div>
  );
}

export default DisplayRecipe;
