package graph_server

import "net/url"

type Exchange struct {
	Method        string
	Path          string
	Query         url.Values
	Form          url.Values
	Authorization string
}
