package repository

import "github.com/proishan11/mux-rest/models"

// PostRepo is interface for posts datastore
type PostRepo interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}
