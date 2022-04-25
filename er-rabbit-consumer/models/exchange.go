package models

import (
	"time"
)

// gorm.Model definition
type Exchanges struct {
	ID   uint `gorm:"primaryKey"`
	Time time.Time
	USD  float64
	EUR  float64
	TRY  float64
}
