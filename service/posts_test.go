package service

import (
	"testing"

	"github.com/proishan11/mux-rest/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock MockRepository) Save(post *models.Post) (*models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*models.Post), args.Error(1)
}

func (mock MockRepository) FindAll() ([]models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]models.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	post := models.Post{ID: 1, Title: "A", Text: "B"}
	// Add expectated results when the given function get called
	mockRepo.On("FindAll").Return([]models.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.FindAll()

	var ID int64 = 1
	mockRepo.AssertExpectations(t)
	assert.Equal(t, ID, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
	assert.Nil(t, err)
}

func TestSave(t *testing.T) {
	mockRepo := new(MockRepository)
	post := models.Post{Title: "A", Text: "B"}
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}

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
