package models

import "time"

type Response struct {
	Result string `json:"result"`
	Time   string `json:"time_last_update_utc"`
	Rate   struct {
		EUR float64 `json:"EUR"`
		TRY float64 `json:"TRY"`
	} `json:"conversion_rates"`
}

type Message struct {
	Time time.Time
	Rate struct {
		EUR float64
		TRY float64
	}
}
