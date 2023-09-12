package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Uid 		uint 	`gorm:"not null"`
	Name 		string 	`gorm:"not null"`
	Done 		bool 	`gorm:"default:false"`
}
