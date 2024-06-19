package httpclient

import (
	"bytes"
	"net/http"
)

type Request struct {
	req *http.Request
}

func NewRequest(method, urlStr string, body []byte) (*Request, error) {
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return &Request{req: req}, nil
}

func (r *Request) SetHeader(key, value string) {
	r.req.Header.Set(key, value)
}

func (r *Request) SetQueryParams(params map[string]string) {
	q := r.req.URL.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	r.req.URL.RawQuery = q.Encode()
}

func (r *Request) GetRequest() *http.Request {
	return r.req
}
