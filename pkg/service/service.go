package service

import "github.com/Doni-githu/golang-test-handler/pkg/repository"

type Person interface {
}

type Service struct {
	Person
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(repos),
	}
}