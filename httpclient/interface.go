package httpclient

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*Response, error)
}

type HTTPRequest interface {
	SetHeader(key, value string)
	SetQueryParams(params map[string]string)
	GetRequest() *http.Request
}
