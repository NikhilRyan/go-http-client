package httpclient

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient(10 * time.Second)
	if client.httpClient.Timeout != 10*time.Second {
		t.Errorf("Expected timeout to be 10s, got %v", client.httpClient.Timeout)
	}
}

func TestDo(t *testing.T) {
	client := NewClient(10 * time.Second)
	req, err := NewRequest(Get, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req.GetRequest())
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func BenchmarkDo(b *testing.B) {
	client := NewClient(10 * time.Second)
	req, err := NewRequest(Get, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		b.Fatalf("Failed to create request: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, err := client.Do(req.GetRequest())
		if err != nil {
			b.Fatalf("Failed to execute request: %v", err)
		}
	}
}
