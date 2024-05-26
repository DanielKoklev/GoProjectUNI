package main

type Book struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Author    string `json:"author"`
    ISBN      string `json:"isbn"`
    Year      int    `json:"year"`
}

var books = []Book{
    {ID: "1", Title: "1984", Author: "George Orwell", ISBN: "9780451524935", Year: 1949},
    {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: "9780061120084", Year: 1960},
}
