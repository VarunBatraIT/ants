package main

import "github.com/ChimeraCoder/anaconda"

type Twitter struct {
	comment string
}

var twitterKey, twitterSecret, twitterToken, twitterTokenSecret string
var api *anaconda.TwitterApi

func init() {
	twitterKey = c.Twitter["key"]
	twitterSecret = c.Twitter["secret"]
	twitterToken = c.Twitter["token"]
	twitterTokenSecret = c.Twitter["tokenSecret"]
	anaconda.SetConsumerKey(twitterKey)
	anaconda.SetConsumerSecret(twitterSecret)
	api = anaconda.NewTwitterApi(twitterToken, twitterTokenSecret)

}

func fixTwitter(t *Twitter) {
	// Add twitter hashes!
	t.comment = AddHashes(Truncate(t.comment, 140), TWITTER_HASH)
}

func PostOnTwitter(t Twitter) {
	fixTwitter(&t)
	api.PostTweet(t.comment, nil)
}
