package failing_dialer

import "net/http"

func (d *Dialer) Client() *http.Client {
	return &http.Client{
		// goanalyze:ignore restricted_literal
		Transport: &http.Transport{DialContext: d.dial}, // goanalyze:ignore restricted_literal
	}
}
