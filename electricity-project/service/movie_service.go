package service

import (
	"fmt"
	"go-practice/electricity-project/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo:repo}
}

func (m *MovieServiceManager) ShowMovieName() string  {
	fmt.Println("MovieName is " + m.repo.GetMovieName())
	return m.repo.GetMovieName()
}
