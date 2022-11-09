package service

import (
	"golang-softplayer-apply/model"
	"golang-softplayer-apply/repository"
)

type Service struct {
	repository repository.RepositoryImpl
}

func NewService(r repository.RepositoryImpl) *Service {
	return &Service{repository: r}
}

func (s *Service) GetAllPeople() []model.Person {
	return s.repository.GetPeople()
}
