package golib

import (
	"io"
	"log"
	"net/http"
)

func PrintHttpResponseBody(resp *http.Response) {
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bodyText)
	Info.Print(s)
}
