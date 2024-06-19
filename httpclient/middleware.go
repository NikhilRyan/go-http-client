package httpclient

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(next http.RoundTripper) http.RoundTripper

type retryTransport struct {
	transport http.RoundTripper
	retries   int
	delay     time.Duration
}

func (rt *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i <= rt.retries; i++ {
		resp, err = rt.transport.RoundTrip(req)
		if err == nil {
			return resp, nil
		}
		log.Printf("Request failed: %v. Retrying... (%d/%d)", err, i, rt.retries)
		time.Sleep(rt.delay)
	}
	return resp, err
}

func RetryMiddleware(retries int, delay time.Duration) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return &retryTransport{
			transport: next,
			retries:   retries,
			delay:     delay,
		}
	}
}

type loggingTransport struct {
	transport http.RoundTripper
}

func (lt *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("Starting request to %s", req.URL)

	resp, err := lt.transport.RoundTrip(req)

	if err != nil {
		log.Printf("Error during request to %s: %v", req.URL, err)
		return nil, err
	}

	log.Printf("Completed request to %s", req.URL)
	log.Printf("Response status: %s", resp.Status)

	return resp, nil
}

func LoggingMiddleware(next http.RoundTripper) http.RoundTripper {
	return &loggingTransport{transport: next}
}
