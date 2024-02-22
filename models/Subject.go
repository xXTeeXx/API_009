// models/Subject.go
package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model 

	Name        string
	Description string
}

