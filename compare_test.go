package main

import (
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func TestCompareAvailabilityUnavailable(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/unavailable.json")

	storeNames := CompareAvailability([]string{"銀座", "心斎橋"}, []string{})
	Test{[]CompareResult{
		CompareResult{"銀座", Deleted},
		CompareResult{"心斎橋", Deleted},
	}, storeNames}.DeepEqual(t)
}

func TestCompareAvailabilityAvailable(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/available.json")

	storeNames := CompareAvailability([]string{"銀座", "心斎橋"}, []string{"銀座", "名古屋栄"})
	Test{[]CompareResult{
		CompareResult{"名古屋栄", Added},
		CompareResult{"心斎橋", Deleted},
	}, storeNames}.DeepEqual(t)
}
