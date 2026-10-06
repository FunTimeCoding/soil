package query

import "net/url"

func New(values url.Values) *Query {
	return &Query{values: values}
}
