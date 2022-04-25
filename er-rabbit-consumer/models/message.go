package models

import "time"

type Message struct {
	Time time.Time
	Rate struct {
		EUR float64
		TRY float64
	}
}
