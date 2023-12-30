package models

import (
	"time"
)

type Person struct {
	ID          int       `gorm:"primarykey"`
	FirstName   string    `gorm:"type:varchar(100)"`
	LastName    string    `gorm:"type:varchar(100)"`
	DateOfBirth time.Time `gorm:"type:date"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   time.Time `gorm:"index"`
}

// TableName overrides the table name used by Person to `person`
func (Person) TableName() string {
	return "person"
}
