package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type LogEntry struct {
	gorm.Model
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}
