# golib
customized golang library for productivity improvement

## Usage

```go
import "github.com/sriram-yeluri/golib"
````

## Example code

```go
c := new(golib.Client)
c.SetBaseURL("https://example.com")
c.SetBasicAuth("username", "password")
resp := c.SendHttpRequest("GET", "https://example.com")
golib.Info.Print(resp.StatusCode)
golib.PrintResponseBody(resp)
```
