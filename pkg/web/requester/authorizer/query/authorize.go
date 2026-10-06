package query

import "net/http"

func (q *Query) Authorize(r *http.Request) error {
	v := r.URL.Query()

	for k, values := range q.values {
		for _, value := range values {
			v.Add(k, value)
		}
	}

	r.URL.RawQuery = v.Encode()

	return nil
}
