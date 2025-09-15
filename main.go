package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// defines and initializes a slice of album structs.
// Each struct holds details about a music album sunch as its id, title, artist and price
// this is kind of like a list of albums structs
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// assign the handler function to an endpoint path
func main() {
	router := gin.Default()
	// when a GET request is made to the /albums endpoint, the getAlbums function will be called
	// the getAlbums function will handle the request and send back a response
	// the response will be a JSON array of all the albums in the albums slice
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// in hw1, you run this router.Run("0.0.0.0:8080") to start the server
	// router.Run("0.0.0.0:8080")

	// in hw2,  you run this router.Run(":8080")
	router.Run(":8080")
}

// getAlbums responds with the list of all albums as JSON.
// getAlbums function that takes gin.Context parameter which carries request details, validates and serialize JSON, and more
// we call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
// the first argument is the HTTP status code, in this case http.StatusOK which is 200
// c stands for context and it reprensents the information about an HTTP request and response. It carries details sunch as the incoming data,
// request headers, and also provides methods to send responses back.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {

	// Declare a variable of type album.
	var newAlbum album

	// c.BindJSON(&newAlbum) will bind the received JSON to newAlbum.
	// It takes data sent by the user in JSON format from the body of the http request and tries to fill
	// the newAlbum variable with this data
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("ID: " + id)

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
