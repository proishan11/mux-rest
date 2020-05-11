package main

import (
	"encoding/json"
	"net/http"
)

// Post is the model for posts
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

// initializes repository
func init() {
	posts = []Post{
		Post{
			ID:    1,
			Title: "Sample title",
			Text:  "Sample Text",
		},
	}
}

// handler function for posts get request
func getPosts(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		// take an integer parameter
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while marshaling posts"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
