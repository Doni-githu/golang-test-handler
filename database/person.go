package database

import "github.com/Doni-githu/golang-test-handler/models"

type Respository struct {
}

func (r *Respository) GetPeople() ([]models.Person, error) {
	return []models.Person{
		{
			Name:    "Doniyor",
			ID:      1,
			Surname: "Doniyorov",
		},
	}, nil
}
