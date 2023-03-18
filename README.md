# golib
customized golang library for productivity improvement

## Usage

```go
import "github.com/sriram-yeluri/golib"
```

### How to print log messages

```go
golib.Info.Print(" This is logged as Info")
golib.Warn.Print(" This is logged as Warning")
golib.Info.Error(" This is logged as Error")
```

### How to make a http request

```go
c := new(golib.Client)
c.SetBaseURL("https://example.com")
c.SetBasicAuth("username", "password")
resp := c.SendHttpRequest("GET", "https://example.com")
```

### General purpose utility functions 

> Print response body golib.PrintResponseBody(resp *http.Response)
