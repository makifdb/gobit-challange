package main

import (
	"er-api-consumer/handlers"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	gc := gocron.NewScheduler(time.UTC)
	gc.Every(1).Minute().Do(func() { handlers.CreateMessage() })
	gc.StartBlocking()
}
