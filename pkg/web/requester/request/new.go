package request

import "net/url"

func New(
	method string,
	path string,
) *Request {
	return &Request{
		Method:     method,
		Path:       path,
		Parameters: url.Values{},
		Header:     map[string]string{},
	}
}
