package repository

import (
	"context"
	"log"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/proishan11/mux-rest/models"
)

type repo struct{}

const (
	projectID      = "mux-rest"
	collectionName = "posts"
)

func (*repo) Save(post *models.Post) (*models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("An error occurred while creating firestore client : %v\n", err)
		return nil, err
	}

	// defer exectues after the function enclosing it returns a value
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("An error occurred while adding posts to datastore : %v\n", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("An error occurred while creating firestore client : %v\n", err)
		return nil, err
	}

	defer client.Close()
	var posts []models.Post
	itr := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("An error occurred while iterating through posts: %v", err)
			return nil, err
		}

		post := models.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// NewFirestoreRepo returns interface for handling firestore data
func NewFirestoreRepo() PostRepo {
	return &repo{}
}
