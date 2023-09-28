package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Merchant struct {
	ID                uint `gorm:"primary_key"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Name              string
	Description       string
	LogoURL           *string
	IsActive          bool
	Email             string
	Phone             string
	UserName          string
	UserEmail         string
	Address           string
	PostalCode        string
	DistrictID        int
	Latitude          float64
	Longitude         float64
	CommissionFee     decimal.Decimal
	BankCode          string
	BankAccountNumber string
}
