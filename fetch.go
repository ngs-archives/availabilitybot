package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// FetchAvailability fetches availability
func FetchAvailability(partNumber string) ([]string, error) {
	v := url.Values{}
	v.Set("parts.0", partNumber)
	v.Set("location", "100-0001")
	resp, err := http.Get("http://www.apple.com/jp/shop/retail/pickup-message?" + v.Encode())
	if err != nil {
		return []string{}, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var obj map[string]interface{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return []string{}, err
	}
	stores := obj["body"].(map[string]interface{})["stores"].([]interface{})
	storeNames := []string{}
	for _, s := range stores {
		store := s.(map[string]interface{})
		quote := store["partsAvailability"].(map[string]interface{})[partNumber].(map[string]interface{})["pickupSearchQuote"].(string)
		if strings.Contains(quote, "本日") {
			storeNames = append(storeNames, store["storeName"].(string))
		}
	}
	return storeNames, nil
}
