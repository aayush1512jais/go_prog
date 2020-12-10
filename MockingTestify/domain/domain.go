package domain

import (
	"github.com/aayush1512jais/go_prog/MockingTestify/repository"
)

type Domain struct {
	appRepo repository.AppRepository
}

func NewDomain(appRepo repository.AppRepository) *Domain {
	return &Domain{
		appRepo: appRepo,
	}
}

//var Err = errors.New("False Error")

type AppDomain interface {
	Get(id int) (int, error)
	Add(id int) error
}

func (db *Domain) GetUser(id int) (int, error) {
	id, err := db.appRepo.Get(id)
	return id, err
}
func (db *Domain) AddUser(id int) error {
	err := db.appRepo.Add(id)
	return err
}
