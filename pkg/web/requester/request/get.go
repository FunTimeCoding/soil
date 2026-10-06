package request

import "net/http"

func Get(path string) *Request {
	return New(http.MethodGet, path)
}
