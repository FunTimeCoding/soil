package unit

import "net/url"

type aptlySeen struct {
	Method string
	Path   string
	Query  url.Values
	Body   string
}
