package main

import (
	"errors"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func TestTweetChanges(t *testing.T) {
	gock.New("https://api.twitter.com").
		Post("/1.1/statuses/update.json").
		Reply(200).
		BodyString("{}")
	gock.New("https://api.twitter.com").
		Post("/1.1/statuses/update.json").
		Reply(200).
		BodyString("{}")

	errors := TweetChanges([]CompareResult{
		CompareResult{"心斎橋", Added},
		CompareResult{"名古屋栄", Deleted},
	})
	Test{[]error{}, errors}.DeepEqual(t)
}

func TestTweetChangesError(t *testing.T) {
	gock.New("https://api.twitter.com").
		Post("/1.1/statuses/update.json").
		ReplyError(errors.New("omg 1"))
	gock.New("https://api.twitter.com").
		Post("/1.1/statuses/update.json").
		ReplyError(errors.New("omg 2"))

	res := TweetChanges([]CompareResult{
		CompareResult{"心斎橋", Added},
		CompareResult{"名古屋栄", Deleted},
	})
	for _, test := range []Test{
		Test{2, len(res)},
		Test{"Post https://api.twitter.com/1.1/statuses/update.json: omg 1", res[0].Error()},
		Test{"Post https://api.twitter.com/1.1/statuses/update.json: omg 2", res[1].Error()},
	} {
		test.Compare(t)
	}
}
