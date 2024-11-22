package service

import (
	"mrs-go/model"
	"mrs-go/repository"

	"gorm.io/gorm"
)

type showDataService interface {
	GetLandmarkByName(name string) (*[]model.Landmark, error)
	UpdatelandmarkByNameAndByDetail(idx int, name string, detail string) *gorm.DB
}

type showDate struct {
	db *gorm.DB
}

func NewshowDataService(gormdb *gorm.DB) showDataService {
	return showDate{db: gormdb}
}

func (s showDate) GetLandmarkByName(name string) (*[]model.Landmark, error) {
	landmarkRepo := repository.NewLandmarkRepo(s.db)
	landmarks, err := landmarkRepo.GetAllLandmarkByName(name)
	if err != nil {
		return nil, err
	}
	return landmarks, nil
}
func (s showDate) UpdatelandmarkByNameAndByDetail(idx int, name string, detail string) *gorm.DB {
	landmarkRepo := repository.NewLandmarkRepo(s.db)
	result := landmarkRepo.PutlandmarkByNameAndByDetail(idx, name, detail)
	if result.Error != nil {
		panic(result.Error)
	}
	return result
}
