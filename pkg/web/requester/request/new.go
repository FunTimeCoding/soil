package request

import (
	"context"
	"net/url"
)

func New(
	method string,
	path string,
) *Request {
	return &Request{
		Context:    context.Background(),
		Method:     method,
		Path:       path,
		Parameters: url.Values{},
		Header:     map[string]string{},
	}
}
