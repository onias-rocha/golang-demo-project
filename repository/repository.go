package repository

import (
	"github.com/jmoiron/sqlx"
	"golang-softplayer-apply/model"
	"time"
)

type RepositoryImpl struct {
	db *sqlx.DB
}

type Repository interface {
	GetPeople() []model.Person
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) GetPeople() []model.Person {
	return []model.Person{
		{
			ID:        "1",
			Name:      "Teste Silva",
			Gender:    "Fluid",
			Email:     "teste.teste@teste.com",
			BirthDate: time.Date(1995, 03, 30, 0, 0, 0, 0, time.UTC),
			City:      "Testelandia",
			Country:   "Brazil",
			Cpf:       "22255588878",
		},
	}
}
