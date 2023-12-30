package models

import (
	"time"
)

type Person struct {
	ID          int       `gorm:"primarykey"`
	FirstName   string    `gorm:"type:varchar(100)"`
	LastName    string    `gorm:"type:varchar(100)"`
	DateOfBirth time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}
