package service

import "github.com/Doni-githu/golang-test-handler/pkg/repository"

type PersonService struct {
	r *repository.Repository
}

func NewPersonService(repo *repository.Repository) *PersonService {
	return &PersonService{r: repo}
}