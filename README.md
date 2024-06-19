
# Go HTTP Client

A Golang library for making HTTP requests with support for middleware, logging, retries, and more.

## Features

- Supports all HTTP methods (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS)
- Middleware support (logging, retries)
- Customizable transport
- JSON and XML parsing utilities
- Logging utilities

## Installation

To install the library, use:

```sh
go get github.com/nikhilryan/go-http-client
```

## Usage

### Basic Usage

Here's an example of how to use the library to make a simple GET request:

```go
package main

import (
    "fmt"
    "time"
    "go-http-client/httpclient"
    "go-http-client/utils"
)

func main() {
    client := httpclient.NewClient(
        10*time.Second,
        httpclient.WithTransport(httpclient.NewTransport(http.DefaultTransport)),
        httpclient.WithMiddleware(httpclient.LoggingMiddleware),
        httpclient.WithMiddleware(httpclient.RetryMiddleware(3, 2*time.Second)),
    )

    req, err := httpclient.NewRequest(httpclient.Get, "https://jsonplaceholder.typicode.com/todos/1", nil)
    if err != nil {
        utils.Error("Error creating request:", err)
        return
    }

    req.SetHeader("Accept", "application/json")

    resp, err := client.Do(req.GetRequest())
    if err != nil {
        utils.Error("Error making request:", err)
        return
    }

    fmt.Println("Status Code:", resp.StatusCode)
    fmt.Println("Response Body:", resp.String())

    var result map[string]interface{}
    if err := resp.ParseJSON(&result); err != nil {
        utils.Error("Error parsing JSON:", err)
        return
    }
    fmt.Println("Parsed Response:", result)
}
```

### Middleware

You can add middleware to your client for logging and retries:

```go
client := httpclient.NewClient(
    10*time.Second,
    httpclient.WithTransport(httpclient.NewTransport(http.DefaultTransport)),
    httpclient.WithMiddleware(httpclient.LoggingMiddleware),
    httpclient.WithMiddleware(httpclient.RetryMiddleware(3, 2*time.Second)),
)
```

### JSON and XML Parsing

Use the provided utilities for JSON and XML parsing:

```go
import (
    "go-http-client/utils"
)

// JSON
data := map[string]interface{}{"key": "value"}
jsonData, err := utils.ToJSON(data)
if err != nil {
    utils.Error("Error marshaling JSON:", err)
}

// XML
type Example struct {
    Key string `xml:"key"`
}
xmlData, err := utils.ToXML(Example{Key: "value"})
if err != nil {
    utils.Error("Error marshaling XML:", err)
}
```

## License

This project is licensed under the MIT License.
