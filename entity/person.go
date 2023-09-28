package entity

import (
	"time"
)

type Person struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Country   string
}

func (Person) TableName() string {
	return "persons"
}
