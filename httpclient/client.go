package httpclient

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient  *http.Client
	middlewares []Middleware
}

type ClientOption func(*Client)

func NewClient(timeout time.Duration, opts ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	// Wrap transport with middlewares
	transport := client.httpClient.Transport
	for _, mw := range client.middlewares {
		transport = mw(transport)
	}
	client.httpClient.Transport = transport

	return client
}

func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return NewResponse(resp)
}

func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *Client) {
		c.httpClient.Transport = transport
	}
}

func WithMiddleware(middleware Middleware) ClientOption {
	return func(c *Client) {
		c.middlewares = append(c.middlewares, middleware)
	}
}
