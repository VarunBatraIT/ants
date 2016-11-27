package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func scheduler() {
	// Refresh our list of blogs and other settings.
	location, err := time.LoadLocation("Asia/Kolkata")
	var c *cron.Cron
	if err != nil {
		log.Println("Unfortunately can't load a location")
		log.Println(err)
		c = cron.New()
	} else {
		log.Println(location)
		c = cron.NewWithLocation(location)
	}
//	c.AddFunc("* * * * * *",func() {
//	})
	c.Start()
}
