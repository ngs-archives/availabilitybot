package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

func run() error {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return nil
	}
	partNumber := flag.Arg(0)
	rc, _ := homedir.Expand("~/.avaiabilitybotrc")
	var currentStoreNames []Store
	rcContent, _ := ioutil.ReadFile(rc)
	json.Unmarshal(rcContent, &currentStoreNames)
	newStoreNames, err := FetchAvailability(partNumber)
	if err != nil {
		return err
	}
	results := CompareAvailability(currentStoreNames, newStoreNames)
	errs := TweetChanges(results)
	for _, err := range errs {
		log.Println(err.Error())
	}
	if len(errs) > 0 {
		return errors.New("Tweet Error")
	}
	data, _ := json.Marshal(&newStoreNames)
	return ioutil.WriteFile(rc, data, os.ModePerm)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
