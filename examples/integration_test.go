package main

import (
	"go-http-client/httpclient"
	"net/http"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	client := httpclient.NewClient(
		10*time.Second,
		httpclient.WithTransport(httpclient.NewTransport(http.DefaultTransport)),
		httpclient.WithMiddleware(httpclient.LoggingMiddleware),
		httpclient.WithMiddleware(httpclient.RetryMiddleware(3, 2*time.Second)),
	)

	req, err := httpclient.NewRequest(httpclient.Get, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	req.SetHeader("Accept", "application/json")

	resp, err := client.Do(req.GetRequest())
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := resp.ParseJSON(&result); err != nil {
		t.Fatalf("Error parsing JSON: %v", err)
	}

	if result["id"].(float64) != 1 {
		t.Errorf("Expected ID 1, got %v", result["id"])
	}
}
