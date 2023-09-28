package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type GetUsersFilter struct {
	Name        string `query:"name"`
	NIK         string `query:"name"`
	ID          *int   `query:"id"`
	PageRequest PageRequest
}

type GetUsersResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}

type GetUserByIDRequest struct {
	UserID int `param:"personId" validate:"required"`
}

type GetUserByIDResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}

type GetUserBalanceByNumber struct {
	Number string `param:"no_rekening" validate:"required"`
}

type GetUserBalanceResult struct {
	Balance decimal.Decimal `json:"saldo"`
}

type EditUserRequest struct {
	UserID int `param:"personId" validate:"required"` // Path variable

	Name    string `json:"name" validate:"required,notblank,min=3,max=60"`
	Country string `json:"country" validate:"required,notblank,min=5,max=250"`
}

type EditUserResult struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
}

type CreateUserRequest struct {
	Name   string `json:"nama" validate:"required,notblank,min=3,max=60"`
	NIK    string `json:"nik" validate:"required,notblank,min=16,max=16"`
	Phone  string `json:"no_hp" validate:"required,notblank,min=11,max=13"`
	Number string `json:"number"`
}

type CreateUserResult struct {
	Name   string `json:"nama"`
	Number string `json:"no_rekening"`
}

type SaveBalanceUserRequest struct {
	Number string `json:"no_rekening" validate:"required,notblank"`
	Amount int    `json:"nominal" validate:"required,notblank"`
}
