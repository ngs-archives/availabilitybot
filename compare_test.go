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

	stores := CompareAvailability([]Store{
		{Name: "銀座", Product: "AirPods"},
		{Name: "心斎橋", Product: "AirPods"},
	}, []Store{})
	Test{[]CompareResult{
		{Store{Name: "銀座", Product: "AirPods"}, Deleted},
		{Store{Name: "心斎橋", Product: "AirPods"}, Deleted},
	}, stores}.DeepEqual(t)
}

func TestCompareAvailabilityAvailable(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/available.json")

	stores := CompareAvailability([]Store{
		{Name: "銀座", Product: "AirPods"},
		{Name: "心斎橋", Product: "AirPods"},
	}, []Store{
		{Name: "銀座", Product: "AirPods"},
		{Name: "名古屋栄", Product: "AirPods"},
	})
	Test{[]CompareResult{
		{Store{Name: "名古屋栄", Product: "AirPods"}, Added},
		{Store{Name: "心斎橋", Product: "AirPods"}, Deleted},
	}, stores}.DeepEqual(t)
}
