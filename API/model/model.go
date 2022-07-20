package model

import (
	"time"
)

type QueueModel struct {
	Code   int    `gorm:"size:5"`
	Type   string `gorm:"size:2"`
	Date   time.Time
	Name   string `gorm:"size:30"`
	Tel    string `gorm:"size:16"`
	UserID string
}

type QueueInput struct {
	Type   string `gorm:"size:2"`
	Name   string `gorm:"size:30"`
	Tel    string `gorm:"size:16"`
	UserID string `gorm:"size:50"`
}

type QueueResponse struct {
	Code string `gorm:"size:5"`
	Date time.Time
	Name string `gorm:"size:30"`
	Tel  string `gorm:"size:16"`
}

type QueuesResponse struct {
	Code string `gorm:"size:5"`
	Type string `gorm:"size:2"`
	Date time.Time
	Name string `gorm:"size:30"`
	Tel  string `gorm:"size:16"`
}

type QueueResponseLine struct {
	CurrentCode string `gorm:"size:5"`
	UserCode    string `gorm:"size:5"`
	QueueAmount int    `gorm:"size:5"`
	Date        time.Time
	Name        string `gorm:"size:30"`
}
