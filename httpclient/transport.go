package httpclient

import (
	"log"
	"net/http"
	"time"
)

type Transport struct {
	transport http.RoundTripper
}

func NewTransport(transport http.RoundTripper) *Transport {
	return &Transport{transport: transport}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	log.Printf("Starting request to %s", req.URL)

	resp, err := t.transport.RoundTrip(req)

	if err != nil {
		log.Printf("Error during request to %s: %v", req.URL, err)
		return nil, err
	}

	log.Printf("Completed request to %s in %v", req.URL, time.Since(start))
	log.Printf("Response status: %s", resp.Status)

	return resp, nil
}
