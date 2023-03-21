package golib

import (
	"io"
	"net/http"
)

func PrintHttpResponseBody(resp *http.Response) error {
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s := string(bodyText)
	Info.Print(s)
	return nil
}
