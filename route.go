package main

import (
	"encoding/json"
	"log"
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
		log.Println("Error occured while marshalling posts")
		// take an integer parameter
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while marshalling posts"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		log.Println("Error occurred while unmarshalling the request")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while unmarshalling the request"}`))
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(post)
	if err != nil {
		log.Println("Error occurred while marshalling the response")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while marshalling the request"}`))
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
