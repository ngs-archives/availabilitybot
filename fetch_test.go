package main

import (
	"errors"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func TestFetchAvailabilityUnavailable(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/unavailable.json")

	storeNames, err := FetchAvailability("MMEF2J/A")
	if err != nil {
		t.Errorf("Expected nil but got '%v'", err.Error())
	}
	Test{[]string{}, storeNames}.DeepEqual(t)
}

func TestFetchAvailabilityAvailable(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/available.json")

	storeNames, err := FetchAvailability("MMEF2J/A")
	if err != nil {
		t.Errorf("Expected nil but got '%v'", err.Error())
	}
	Test{[]string{"銀座", "名古屋栄"}, storeNames}.DeepEqual(t)
}

func TestFetchAvailabilityHTTPError(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		ReplyError(errors.New("omg"))

	storeNames, err := FetchAvailability("MMEF2J/A")
	Test{
		"Get http://www.apple.com/jp/shop/retail/pickup-message?location=100-0001&parts.0=MMEF2J%2FA: omg",
		err.Error(),
	}.Compare(t)
	Test{[]string{}, storeNames}.DeepEqual(t)
}

func TestFetchAvailabilityInvalidJSON(t *testing.T) {
	defer gock.Off()
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		BodyString("{invalid}")

	storeNames, err := FetchAvailability("MMEF2J/A")
	Test{
		"invalid character 'i' looking for beginning of object key string",
		err.Error(),
	}.Compare(t)
	Test{[]string{}, storeNames}.DeepEqual(t)
}
