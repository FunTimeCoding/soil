package request

import "net/url"

func (q *Request) WithParameters(v url.Values) *Request {
	for k, list := range v {
		for _, value := range list {
			q.Parameters.Add(k, value)
		}
	}

	return q
}
