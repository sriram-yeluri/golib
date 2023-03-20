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
golib.Error.Print(" This is logged as Error")
```

### How to make a http request

```go
c := new(golib.Client)
c.SetBaseURL("https://example.com")
c.SetBasicAuth("username", "password")
resp := c.SendHttpRequest("GET", "/api/v1/endpoint")
```

### General purpose utility functions 

Print response body : `golib.PrintHttpResponseBody(resp)(resp *http.Response)`  
Write HttpResponse to a given file name : `golib.WriteHttpResponseToFile(resp, "resp.txt")`  
Read from a given file and return []byte data : `data := golib.ReadFromFile("resp.txt")`  



