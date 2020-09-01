package repositories

import "go-practice/electricity-project/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name: "totti"}
	return movie.Name
}
