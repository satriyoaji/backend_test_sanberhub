package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Mutation struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Number    string
	Code      string
	Amount    decimal.Decimal
}
