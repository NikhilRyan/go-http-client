package httpclient

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func NewResponse(resp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}

func (r *Response) ParseJSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

func (r *Response) ParseXML(v interface{}) error {
	return xml.Unmarshal(r.Body, v)
}

func (r *Response) String() string {
	return string(r.Body)
}
