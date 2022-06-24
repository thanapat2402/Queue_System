package model

import "time"

type QueueModel struct {
	Date      time.Time
	TableType string
	QueueNo   int
}
