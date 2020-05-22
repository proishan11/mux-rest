package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/proishan11/mux-rest/errors"

	"github.com/proishan11/mux-rest/models"
	"github.com/proishan11/mux-rest/service"
)

var (
	postService service.PostService
)

type postController struct{}

type PostController interface {
	GetPosts(http.ResponseWriter, *http.Request)
	AddPost(http.ResponseWriter, *http.Request)
}

// handler function for posts get request
func (*postController) GetPosts(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		log.Println("Error occurred while fetching the post")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error occurred while fetching the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*postController) AddPost(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post models.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		log.Println("Error occurred while unmarshalling the request")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error occurred while unmarshalling the request"})
		return
	}

	err = postService.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := postService.Create(&post)

	if err != nil {
		log.Println("Error occurred while adding the post")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error occurred while saving the post"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &postController{}
}
