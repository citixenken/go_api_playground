package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

/*store book data in memory => struct tags ie `json:"title"`
specify shape when content is serialized into JSON */
type book struct {
	// capitalization makes it an exported field/public field
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

// books slice to seed record book data.
var books = []book{
	{ID: "1", Title: "Dune", Author: "Frank Herbert", Quantity: 5},
    {ID: "2", Title: "Fahrenheit 451", Author: "Ray Bradbury", Quantity: 3},
    {ID: "3", Title: "Hail Mary", Author: "Andy Weir", Quantity: 14},
}

func main (){
	router := gin.Default()
	router.GET("/books", getBooks)
	// router.GET("/books/:id", getBookByID)
	router.GET("/books/:id", getBook)
	router.POST("/books", postBook)
	router.PATCH("/checkout", checkoutBook)

	router.Run("localhost:8080") //attach the router to an http.Server and start the server
}

// getBooks responds with the list of all books as JSON.
//*gin.Context => stores info relating to a specific request
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

// postBook adds a book from JSON received in the request body.
func postBook(c *gin.Context){
	var newBook book

	// Call BindJSON to bind the received JSON to
    // newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
// func getBookByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of books, looking for
//     // a book whose ID value matches the parameter.
// 	for _, b := range books {
// 		if b.ID == id {
// 			c.IndentedJSON(http.StatusOK, b)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"book not found"})
// }

// alternatively;
// 1. helper function
func getBookByID(id string) (*book, error) {
	for i,b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// 2. Actual getBook code
func getBook(c *gin.Context){
	//path param
	id := c.Param("id")

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context){
	// query param => e.g. checkout?id=2
	id, ok := c.GetQuery("id")

	if !ok  {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	// allow checkout only when book.Quantity > 0
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	// checkout book
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}