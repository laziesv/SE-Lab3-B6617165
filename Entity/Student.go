package entity

import (
	
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	FirstName string `valid:"required~First name is required"`
	LastName  string `valid:"required~Last name is required"`

	Email string `valid:"required~Email is required,email~Email format is invalid"`

	Age uint `valid:"required~Age is required,range(1|150)~Age must be between 1 and 150"`
}



