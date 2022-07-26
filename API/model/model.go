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
	Code   string `gorm:"size:5"`
	Date   time.Time
	Name   string `gorm:"size:30"`
	Tel    string `gorm:"size:16"`
	UserID string `gorm:"size:50"`
}

type QueuesResponse struct {
	Code string `gorm:"size:5"`
	Type string `gorm:"size:2"`
	Date time.Time
	Name string `gorm:"size:30"`
	Tel  string `gorm:"size:16"`
}

type QueueResponseLine struct {
	UserCode    string `gorm:"size:5"`
	QueueAmount string `gorm:"size:15"`
	Date        time.Time
	Name        string `gorm:"size:30"`
}

type ReportQueue struct {
	AmountQueueA  string `gorm:"size:5"`
	AmountQueueB  string `gorm:"size:5"`
	AmountQueueC  string `gorm:"size:5"`
	AmountQueueD  string `gorm:"size:5"`
	CurrentQueueA string `gorm:"size:2"`
	CurrentQueueB string `gorm:"size:2"`
	CurrentQueueC string `gorm:"size:2"`
	CurrentQueueD string `gorm:"size:2"`
}

type LineProfile struct {
	UserID      string
	DisplayName string
	PicUrl      string
	Status      string
	Language    string
}

