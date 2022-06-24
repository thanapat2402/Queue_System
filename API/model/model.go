package model

import (
	"time"
)

type QueueModel struct {
	Code   int    `gorm:"size:5"`
	Type   string `gorm:"size:2"`
	Date   time.Time
	Status string `gorm:"size:20;default:Queue"`
}

type QueueOp struct {
	Code string
	Type string
	Date time.Time
}

// func (Q QueueModel) StrCode() string {
// 	return fmt.Sprintf("%v%03d", Q.Type, Q.Code)
// }
