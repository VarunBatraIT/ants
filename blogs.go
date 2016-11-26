package main

import (
	"github.com/mmcdole/gofeed"
	"github.com/parnurzeal/gorequest"
)

func ReadBlogFeed(rss_uri string) (*gofeed.Feed, error) {
	request := gorequest.New()
	_, feedData, err := request.Get(rss_uri).
		Set("User-Agent", o.UserAgents.sample()).
		End()
	var feed *gofeed.Feed
	if err != nil {
		return feed, err[0]
	}
	fp := gofeed.NewParser()
	var parseError error
	feed, parseError = fp.ParseString(feedData)
	return feed, parseError
}

func GetBlogFeed(typeOfBlog string) string {
	if val, ok := o.Blogs[typeOfBlog]; ok {
		return SampleString(val)
	}else{
		return o.Blogs.sample()
	}
}
