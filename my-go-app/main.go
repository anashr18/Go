package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func ListRecipeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"recipes": recipes,
	})
}

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"title": name, "subtitle": "Greetings, ",
		"message": "Welcome to the Index Page!", "version": "1.0",
		"date": "2022-10-20", "author": "John Doe",
		"description": "This is a sample API response", "status": "success",
	})
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	// c.JSON(
	// 	http.StatusOK,
	// 	recipe,
	// )
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Recipe added successfully", "recipe": recipe})
}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedRecipe Recipe
	if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, r := range recipes {
		if r.ID == id {
			recipes[i] = updatedRecipe
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Recipe updated successfully", "recipe": updatedRecipe})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.Run(":8000")
}
