package basic_transport

import "net/http"

func (t *Transport) RoundTrip(e *http.Request) (*http.Response, error) {
	e.SetBasicAuth(t.username, t.password)

	return t.r.RoundTrip(e)
}
