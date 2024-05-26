package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/books", getBooks)
    r.GET("/book/:id", getBook)
    r.POST("/book", createBook)
    r.PUT("/book/:id", updateBook)
    r.DELETE("/book/:id", deleteBook)

    r.Run()
}
