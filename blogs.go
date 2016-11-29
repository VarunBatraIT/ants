package main

import (
	"reflect"
	"github.com/mmcdole/gofeed"
	"github.com/parnurzeal/gorequest"
)

func ReadBlogFeed(rss_uri string) (*gofeed.Feed, error) {
	var parseError error
	var feed *gofeed.Feed
	request := gorequest.New()
	_, feedData, err := request.Get(rss_uri).
		Set("User-Agent", o.UserAgents.sample()).
		End()
	if err != nil {
		return feed, err[0]
	}
	fp := gofeed.NewParser()
	feed, parseError = fp.ParseString(feedData)
	return feed, parseError
}

func GetBlogFeed(typeOfBlog string) string {
	if val, ok := o.Blogs[typeOfBlog]; ok {
		blog := SampleString(val)
		return reflect.ValueOf(blog).String()

	} else {
		return o.Blogs.sample()
	}
}

func GetBlogFeedItem(feed *gofeed.Feed) *gofeed.Item {
	return SampleString(feed.Items).(*gofeed.Item)
}

func PostBlogOnLinkedin(typeOfBlog string) {
	rss_uri := GetBlogFeed(typeOfBlog)
	feed, err := ReadBlogFeed(rss_uri)
	item := GetBlogFeedItem(feed)
	if err == nil {
		toPost := Linkedin{comment: item.Title + " " + item.Link}
		PostOnLinkedin(toPost)
	}

}
