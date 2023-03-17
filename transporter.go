package golib

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func SendHttpRequest(httpMethod string, auth *Auth, requestURL string) (*http.Response, string) {

	//validate the authentication credentials
	if auth.UserName == "" || auth.Password == "" {
		Error.Print("Invalid Credentials")
		os.Exit(1)

	}
	// For control over proxies, TLS configuration, keep-alives,
	// compression, and other settings, create a Transport:
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	// Set http client
	client := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}
	// Create http request
	req, err := http.NewRequest(httpMethod, requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set http request headers
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(auth.UserName, auth.Password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bodyText)

	// return http response and response body as text
	return resp, s

}
