package main

import (
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
}