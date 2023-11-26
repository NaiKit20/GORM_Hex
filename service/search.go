package service

import (
	"gorm-hex/model"
	"gorm-hex/repository"
)

// service struct (<name>Serv)
type searchServ struct{}

// service contructor (New<name>Service)
func NewSearchService() SearchService {
	return searchServ{}
}

// create repositories
var landmarkRepo = repository.NewLandmarkRepository()

// service interface (<name>Service)
type SearchService interface {
	GetAllLandmarks() ([]model.Landmark, error)
	FindByName(name string) ([]model.Landmark, error)
}

func (d searchServ) GetAllLandmarks() ([]model.Landmark, error) {

	landmarks, err := landmarkRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return landmarks, nil
}

func (d searchServ) FindByName(name string) ([]model.Landmark, error) {

	landmarks, err := landmarkRepo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return landmarks, nil
}

