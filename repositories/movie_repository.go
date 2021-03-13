package repositories

import (
	"go-mvc-demo/datamodels"
)

type MovieReposity interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieReposity {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	// 模拟从数据库查询值赋值给模型
	movies :=&datamodels.Movie{Name:"mvc demo"}
	return movies.Name
}