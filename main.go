package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json."id"`
	Title  string `json."title"`
	Author string `json."author"`
	Pages  int    `json."pages"`
}

var books = []book{
	{ID: "1", Title: "Things Fall Apart", Author: "Chinua Achebe", Pages: 209},
	{ID: "2", Title: "Fairy Tales", Author: "Hans Christian Andersen", Pages: 784},
	{ID: "2", Title: "The Divine Comedy", Author: "Dante Alighieri", Pages: 928},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
