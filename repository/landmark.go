package repository

import (
	"mrs-go/model"

	"gorm.io/gorm"
)

type LandmarkRepo interface {
	GetAllLandmarkByName(name string) (*[]model.Landmark, error)
	PutlandmarkByNameAndByDetail(idx int, name string, detail string) *gorm.DB
}
type landmarkDB struct {
	db *gorm.DB
}

func NewLandmarkRepo(gormdb *gorm.DB) LandmarkRepo {
	return landmarkDB{db: gormdb}
}

func (l landmarkDB) GetAllLandmarkByName(name string) (*[]model.Landmark, error) {
	landmarks := []model.Landmark{}
	name = "%" + name + "%"
	result := l.db.Preload("Country").Where("Name like ?", name).Find(&landmarks)
	if result.Error != nil {
		return nil, result.Error
	}
	return &landmarks, nil
}

func (l landmarkDB) PutlandmarkByNameAndByDetail(idx int, name string, detail string) *gorm.DB {
	result := l.db.Model(model.Landmark{}).Where("idx =?", idx).Updates(model.Landmark{Name: name, Detail: detail})
	if result.Error != nil {
		panic(result.Error)
	}
	return result
}
