package model

import (
	"time"
)

type QueueModel struct {
	Code string `gorm:"size:5"`
	Type string `gorm:"size:2"`
	Date time.Time
	Name string `gorm:"size:30"`
	Tel  string `gorm:"size:16"`
	// Status string `gorm:"size:20;default:Queue"`
}

type QueueInput struct {
	Type string `gorm:"size:2"`
	Name string `gorm:"size:30"`
	Tel  string `gorm:"size:16"`
}
