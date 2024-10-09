package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
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
func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.Run(":8000")
}
