package golib

import (
	"testing"
)

func TestPrintHttpResponseBody(t *testing.T) {
	var c = new(Client)
	c.SetBaseURL("http://ip.jsontest.com/")
	c.SetBasicAuth("username", "password")
	resp, err := c.SendHttpRequest("GET", "")
	if err != nil {
		Error.Print(err)
	}
	PrintHttpResponseBody(resp)
}
