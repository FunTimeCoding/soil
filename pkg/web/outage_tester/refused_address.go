package outage_tester

import (
	"net/http"
	"net/http/httptest"
)

func (o *Tester) RefusedAddress() string {
	o.t.Helper()
	s := httptest.NewServer(http.NotFoundHandler())
	s.Close()

	return s.URL
}
