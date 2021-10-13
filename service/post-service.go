package service

import (
	"errors"
	"math/rand"

	"github.com/kristianespina/golang-rest-api/entity"
	"github.com/kristianespina/golang-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

// NewPostService creates new PostService
func NewPostService(newRepo repository.PostRepository) PostService {
	repo = newRepo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	// Check if post exists
	if post == nil {
		err := errors.New("Post cannot be empty")
		return err
	}

	// Check if title is not empty
	if post.Title == "" {
		err := errors.New("Post title is empty")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
