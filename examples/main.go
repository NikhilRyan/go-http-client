package main

import (
	"fmt"
	"go-http-client/httpclient"
	"go-http-client/utils"
	"net/http"
	"time"
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
