package golib

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type User struct {
	UserName string
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
	c.UserInfo.UserName = username
	c.UserInfo.Password = password
	return c
}

func (c *Client) SendHttpRequest(httpMethod string, requestURL string) (*http.Response) {

	//validate the authentication credentials
	if c.UserInfo.UserName == "" || c.UserInfo.Password == "" {
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
	req.SetBasicAuth(c.UserInfo.UserName, c.UserInfo.Password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// return http response
	return resp
}

func PrintResponseBody(resp *http.Response){
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bodyText)
	Info.Print(s)
}
