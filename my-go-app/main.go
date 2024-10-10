package main

import (
	"net/http"
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

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run(":8080")
}

// func main() {
// 	router := gin.Default()
// 	router.GET("/:name", IndexHandler)
// 	router.Run(":8080")
// }

// {
// 	"name": "Homemade Pizza",
// 	"tags" : ["italian", "pizza", "dinner"],
// 	"ingredients": [
// 	"1 1/2 cups (355 ml) warm water (105°F-115°F)",
// 	"1 package (2 1/4 teaspoons) of active dry yeast",
// 	"3 3/4 cups (490 g) bread flour",
// 	"feta cheese, firm mozzarella cheese, grated"
// 	],
// 	"instructions": [
// 		"Step 1.",
// 	"Step 2.",
// 	"Step 3."
// 	]
// }
