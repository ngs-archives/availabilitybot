package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
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

func setup(t *testing.T) {
	rc, _ := homedir.Expand("~/.avaiabilitybotrc")
	os.Remove(rc)
	if _, err := os.Stat(rc); err == nil {
		t.Fatalf("File exists at %v", rc)
	}
}

func readrc() string {
	rc, _ := homedir.Expand("~/.avaiabilitybotrc")
	b, _ := ioutil.ReadFile(rc)
	return string(b)
}

func TestNoArg(t *testing.T) {
	setup(t)
	statuses := []string{}
	setUpdateStatusFn(func(c *twitter.Client, status string) error {
		statuses = append(statuses, status)
		return nil
	})
	os.Args = []string{"app"}
	err := run()

	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	Test{[]string{}, statuses}.DeepEqual(t)
}

func TestAvailable(t *testing.T) {
	defer gock.Off()
	setup(t)
	statuses := []string{}
	setUpdateStatusFn(func(c *twitter.Client, status string) error {
		statuses = append(statuses, status)
		return nil
	})
	os.Args = []string{"app", "MMEF2J/A"}
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/available.json")
	err := run()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	Test{
		[]string{
			"銀座店に AirPods の在庫が追加されました",
			"名古屋栄店に AirPods の在庫が追加されました",
		},
		statuses,
	}.DeepEqual(t)

	Test{`[{"Name":"銀座","Product":"AirPods"},{"Name":"名古屋栄","Product":"AirPods"}]`, readrc()}.Compare(t)
}

func TestTweetError(t *testing.T) {
	defer gock.Off()
	setup(t)
	setUpdateStatusFn(func(c *twitter.Client, status string) error {
		return fmt.Errorf("omg %v", status)
	})
	os.Args = []string{"app", "MMEF2J/A"}
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		Reply(200).
		File("_fixtures/available.json")
	err := run()
	Test{"Tweet Error", err.Error()}.Compare(t)
	Test{"", readrc()}.Compare(t)
}

func TestJSONError(t *testing.T) {
	defer gock.Off()
	setup(t)
	statuses := []string{}
	setUpdateStatusFn(func(c *twitter.Client, status string) error {
		statuses = append(statuses, status)
		return nil
	})

	os.Args = []string{"app", "MMEF2J/A"}

	statuses = []string{}
	gock.New("http://www.apple.com").
		Get("/jp/shop/retail/pickup-message").
		ReplyError(errors.New("omg"))

	err := run()
	Test{"Get http://www.apple.com/jp/shop/retail/pickup-message?location=100-0001&parts.0=MMEF2J%2FA: omg", err.Error()}.Compare(t)

	Test{
		[]string{},
		statuses,
	}.DeepEqual(t)

	Test{"", readrc()}.Compare(t)

}
