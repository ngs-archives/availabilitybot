package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}
	partNumber := flag.Args()[0]
	rc, _ := homedir.Expand("~/.avaiabilitybotrc")
	var currentStoreNames []string
	rcContent, _ := ioutil.ReadFile(rc)
	json.Unmarshal(rcContent, &currentStoreNames)
	newStoreNames, err := FetchAvailability(partNumber)
	if err != nil {
		log.Fatal(err)
	}
	results := CompareAvailability(currentStoreNames, newStoreNames)
	errs := TweetChanges(results)
	for _, err := range errs {
		log.Println(err.Error())
	}
	if len(errs) > 0 {
		log.Fatal("Tweet Error")
	}
	data, err := json.Marshal(&newStoreNames)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(rc, data, os.ModePerm)
}
