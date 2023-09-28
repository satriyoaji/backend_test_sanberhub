package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uint `gorm:"primary_key"`
	MerchantID  uint
	Name        string
	Description string
	Qty         int
	Price       decimal.Decimal
	Weight      int
	Length      int
	Width       int
	Height      int
	ImageURL    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
