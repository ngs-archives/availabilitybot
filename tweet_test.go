package main

import (
	"fmt"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
)

func TestTweetChanges(t *testing.T) {
	statuses := []string{}
	setUpdateStatusFn(func(c *twitter.Client, status string) error {
		statuses = append(statuses, status)
		return nil
	})
	errors := TweetChanges([]CompareResult{
		CompareResult{Store{Name: "心斎橋", Product: "AirPods"}, Added},
		CompareResult{Store{Name: "名古屋栄", Product: "AirPods"}, Deleted},
	})
	Test{[]error{}, errors}.DeepEqual(t)
	Test{[]string{
		"心斎橋店に AirPods の在庫が追加されました",
		"名古屋栄店に AirPods の在庫が無くなりました",
	}, statuses}.DeepEqual(t)
}

func TestTweetChangesError(t *testing.T) {
	setUpdateStatusFn(func(c *twitter.Client, status string) error {
		return fmt.Errorf("omg: %v", status)
	})
	res := TweetChanges([]CompareResult{
		CompareResult{Store{Name: "心斎橋", Product: "AirPods"}, Added},
		CompareResult{Store{Name: "名古屋栄", Product: "AirPods"}, Deleted},
	})
	for _, test := range []Test{
		Test{2, len(res)},
		Test{"omg: 心斎橋店に AirPods の在庫が追加されました", res[0].Error()},
		Test{"omg: 名古屋栄店に AirPods の在庫が無くなりました", res[1].Error()},
	} {
		test.Compare(t)
	}
}
