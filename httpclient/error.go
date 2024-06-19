package httpclient

import "errors"

var (
	ErrInvalidMethod = errors.New("invalid HTTP method")
)
