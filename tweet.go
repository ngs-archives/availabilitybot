package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TwitterClient creates new Twitter client
func TwitterClient() *twitter.Client {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return twitter.NewClient(httpClient)
}

// TweetChanges twees compare results
func TweetChanges(results []CompareResult) []error {
	c := TwitterClient()
	errors := []error{}
	for _, r := range results {
		switch r.Change {
		case Added:
			if _, _, err := c.Statuses.Update(r.StoreName+"店に Airpods の在庫が追加されました", nil); err != nil {
				errors = append(errors, err)
			}
			break
		case Deleted:
			if _, _, err := c.Statuses.Update(r.StoreName+"店に Airpods の在庫が無くなりました", nil); err != nil {
				errors = append(errors, err)
			}
			break
		}
	}
	return errors
}
