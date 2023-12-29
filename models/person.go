package models

import (
	"time"
)

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}
