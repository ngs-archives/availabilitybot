package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	homedir "github.com/mitchellh/go-homedir"

	gock "gopkg.in/h2non/gock.v1"
)

type Test struct {
	expected interface{}
	actual   interface{}
}

func (test Test) Compare(t *testing.T) {
	if test.expected != test.actual {
		t.Errorf(`Expected "%v" but got "%v"`, test.expected, test.actual)
	}
}

func (test Test) DeepEqual(t *testing.T) {
	if !reflect.DeepEqual(test.expected, test.actual) {
		t.Errorf(`Expected "%v" but got "%v"`, test.expected, test.actual)
	}
}

func TestMain(t *testing.T) {
	defer gock.Off()
	rc, _ := homedir.Expand("~/.avaiabilitybotrc")

	os.Remove(rc)
	if _, err := os.Stat(rc); err == nil {
		t.Fatalf("File exists at %v", rc)
	}
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/available.json")
	gock.New("https://api.twitter.com").
		Post("/1.1/statuses/update.json").
		Reply(200).
		BodyString("{}")
	gock.New("https://api.twitter.com").
		Post("/1.1/statuses/update.json").
		Reply(200).
		BodyString("{}")
	os.Args = []string{"test", "MMEF2J/A"}
	main()
	b, _ := ioutil.ReadFile(rc)
	Test{`["銀座","名古屋栄"]`, string(b)}.Compare(t)
}
