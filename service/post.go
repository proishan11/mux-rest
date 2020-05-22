package service

import (
	"errors"
	"math/rand"

	"github.com/proishan11/mux-rest/repository"

	"github.com/proishan11/mux-rest/models"
)

// PostService is interface for post service
type PostService interface {
	Validate(post *models.Post) error
	Create(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}

type postService struct{}

var (
	repo repository.PostRepo = repository.NewFirestoreRepo()
)

func (*postService) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("Empty post")
		return err
	}
	if post.Title == "" {
		err := errors.New("Missing title from post")
		return err
	}

	return nil
}

func (*postService) Create(post *models.Post) (*models.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*postService) FindAll() ([]models.Post, error) {
	return repo.FindAll()
}

// NewPostService returns post service interface
func NewPostService(postRepo repository.PostRepo) PostService {
	repo = postRepo
	return &postService{}
}
