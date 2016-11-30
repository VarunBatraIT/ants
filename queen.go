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
	// Post self marketing!
	c.AddFunc("0 30 8 * * 1,6", func(){
		PostSelfMarketingOnLinkedin()
	})
	// Post jobs!
	c.AddFunc("0 30 8 * * 1,6", func(){
		PostJobOnLinkedin()
	})
	c.AddFunc("0 30 9,19 * * *", func(){
		PostBlogOnLinkedin("leadership")
		PostBlogOnTwitter("leadership")
	})
	c.AddFunc("0 35 9,19 * * *", func(){
		PostBlogOnLinkedin("leadership")
		PostBlogOnTwitter("leadership")
	})
	c.AddFunc("0 45 9,19 * * *", func(){
		PostBlogOnLinkedin("marketing")
		PostBlogOnTwitter("marketing")
	})
	c.AddFunc("0 0 10,20 * * *", func(){
		PostBlogOnLinkedin("technical")
		PostBlogOnTwitter("technical")
	})
	c.Start()
}
