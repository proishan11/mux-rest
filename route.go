package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/proishan11/mux-rest/models"
	"github.com/proishan11/mux-rest/repository"
)

var (
	repo repository.PostRepo = repository.NewPostRepo()
)

// handler function for posts get request
func getPosts(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		log.Println("Error occurred while fetching the post")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while fetching the posts"}`))
	}
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
	var post models.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		log.Println("Error occurred while unmarshalling the request")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while unmarshalling the request"}`))
		return
	}

	// need to fix this
	post.ID = rand.Int63()
	result, err := repo.Save(&post)

	if err != nil {
		log.Println("Error occurred while adding the post")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while adding the post"}`))
		return
	}

	ret, err := json.Marshal(result)
	if err != nil {
		log.Println("Error occurred while marshalling the response")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error occurred while marshalling the request"}`))
	}

	response.WriteHeader(http.StatusOK)
	response.Write(ret)
}
