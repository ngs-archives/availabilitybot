package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Store represents store
type Store struct {
	Name    string
	Product string
}

// FetchAvailability fetches availability
func FetchAvailability(partNumber string) ([]Store, error) {
	v := url.Values{}
	v.Set("parts.0", partNumber)
	v.Set("location", "100-0001")
	resp, err := http.Get("http://www.apple.com/jp/shop/retail/pickup-message?" + v.Encode())
	if err != nil {
		return []Store{}, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var obj map[string]interface{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return []Store{}, err
	}
	storeResults := obj["body"].(map[string]interface{})["stores"].([]interface{})
	stores := []Store{}
	for _, s := range storeResults {
		store := s.(map[string]interface{})
		part := store["partsAvailability"].(map[string]interface{})[partNumber].(map[string]interface{})
		quote := part["pickupSearchQuote"].(string)
		partName := part["storePickupProductTitle"].(string)
		if strings.Contains(quote, "本日") {
			stores = append(stores, Store{
				Name:    store["storeName"].(string),
				Product: partName,
			})
		}
	}
	return stores, nil
}
