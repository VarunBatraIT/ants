package main

import (
	"log"

	"github.com/parnurzeal/gorequest"
)

var linkedinToken string

func init(){
	linkedinToken = c.Linkedin["token"]
}

func PostOnLinkedin(comment, title, description, url, img, code string) {
	uri := "https://api.linkedin.com/v1/people/~/shares?format=json"
	if code == "" {
		code = "anyone"
	}
	var sendString string
	if url != "" {
		sendString = `{
		  "comment" : "` + AddHashes(Truncate(comment, 700), 0) + `",
		  "content": {
		    "title": "` + Truncate(title, 200) + `",
		    "description": "` + Truncate(description, 256) + `",
		    "submitted-url": "` + url + `",
		    "submitted-image-url": "` + img + `"
		  },
		  "visibility": {
		    "code": "` + code + `"
		  }
		}`
	} else {
		sendString = `{
		  "comment" : "` + AddHashes(Truncate(comment, 700), 0) + `",
		  "visibility": {
		    "code": "` + code + `"
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
