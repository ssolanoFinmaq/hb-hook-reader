package entities

import "time"

type ErrorLog struct {
	ID    int
	Time  time.Time `gorm:"default:now()"`
	Error string
}
