package model

import "time"

type GetPersonsFilter struct {
	Name        string `query:"name"`
	IsActive    *bool  `query:"is_active"`
	ID          *int   `query:"id"`
	PageRequest PageRequest
}

type GetPersonsResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}

type CreatePersonRequest struct {
	Name    string `json:"name" validate:"required,notblank,min=3,max=60"`
	Country string `json:"country" validate:"required,notblank,min=5,max=250"`
}

type CreatePersonResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}

type GetPersonByIDRequest struct {
	PersonID int `param:"personId" validate:"required"`
}

type GetPersonCountryByNameRequest struct {
	Name string `param:"name" validate:"required"`
}

type GetPersonByIDResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}

type EditPersonRequest struct {
	PersonID int `param:"personId" validate:"required"` // Path variable

	Name    string `json:"name" validate:"required,notblank,min=3,max=60"`
	Country string `json:"country" validate:"required,notblank,min=5,max=250"`
}

type EditPersonResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}
