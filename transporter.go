package golib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type User struct {
	Username string
	Password string
}

type Client struct {
	UserInfo *User
	BaseURL  string
}

func (c *Client) SetBaseURL(url string) *Client {
	c.BaseURL = strings.TrimRight(url, "/")
	return c
}

func (c *Client) SetBasicAuth(username, password string) *Client {
	c.UserInfo = &User{Username: username, Password: password}
	return c
}

func (c *Client) SendHttpRequest(httpMethod string, api string) *http.Response {

	//validate the authentication credentials
	if c.UserInfo.Username == "" || c.UserInfo.Password == "" {
		Error.Print("Missing Credentials")
		os.Exit(1)
	}

	//form the restEndPoint url
	if c.BaseURL == "" {
		Error.Print("Missing BaseURL")
		os.Exit(1)
	}
	restEndPoint := fmt.Sprintf("%s/%s", c.BaseURL, api)
	if api == "" {
		restEndPoint = strings.TrimRight(restEndPoint, "/")
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
	req, err := http.NewRequest(httpMethod, restEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set http request headers
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.UserInfo.Username, c.UserInfo.Password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// return http response
	return resp
}
