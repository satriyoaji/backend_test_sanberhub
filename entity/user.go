package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Number    string
	Name      string
	NIK       string
	Phone     string
	Balance   decimal.Decimal
}
