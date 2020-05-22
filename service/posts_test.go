package service

import (
	"testing"

	"github.com/proishan11/mux-rest/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Empty post")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := models.Post{ID: 1, Title: "", Text: "Sample Text"}
	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Missing title from post")
}
