package main

import (
	"log"

	"github.com/parnurzeal/gorequest"
)
type Linkedin struct {
	comment string
	title string
	description string
	url string
	img string
	code string
}

var linkedinKey, linkedinSecret, linkedinToken string

func init() {
	linkedinKey = c.Linkedin["key"]
	linkedinSecret = c.Linkedin["secret"]
	linkedinToken = c.Linkedin["token"]
	l := Linkedin{comment: "Testing API"}
	PostOnLinkedin(l)
	log.Println(l)
}

func fixLinkedin(l *Linkedin){
	l.comment = Truncate(l.comment, 700)
	l.title = Truncate(l.title, 200)
	l.description = Truncate(l.description, 256)
}

func PostOnLinkedin(l Linkedin) {
	uri := "https://api.linkedin.com/v1/people/~/shares?format=json"
	if l.code == "" {
		l.code = "anyone"
	}
	fixLinkedin(&l)
	var sendString string
	if l.url != "" {
		sendString = `{
		  "comment" : "` + AddHashes(l.comment, 0) + `",
		  "content": {
		    "title": "` + l.title + `",
		    "description": "` + l.description + `",
		    "submitted-url": "` + l.url + `",
		    "submitted-image-url": "` + l.img + `"
		  },
		  "visibility": {
		    "code": "` + l.code + `"
		  }
		}`
	} else {
		sendString = `{
		  "comment" : "` + AddHashes(l.comment, 0) + `",
		  "visibility": {
		    "code": "` + l.code + `"
		  }
		}`
	}
	log.Println(sendString)
	request := gorequest.New()
	_, body, err := request.Post(uri).
		Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12").
		Set("Content-Type", "application/json").
		Set("x-li-format", "json").
		Set("Authorization", "Bearer "+linkedinToken).
		SendString(sendString).
		End()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(body)
}
