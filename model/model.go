package model

import "time"

type Person struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Cpf       string    `json:"cpf"`
}
