package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
    id := c.Param("id")
    for _, book := range books {
        if book.ID == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func createBook(c *gin.Context) {
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    books = append(books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book
    if err := c.BindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
