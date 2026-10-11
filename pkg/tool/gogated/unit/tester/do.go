package tester

import "net/http"

func (o *Tester) Do(r *http.Request) (*http.Response, error) {
	return o.client.Do(r)
}
