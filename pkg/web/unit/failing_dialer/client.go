package failing_dialer

import "net/http"

func (d *Dialer) Client() *http.Client {
	return &http.Client{ // goanalyze:ignore restricted_literal
		Transport: &http.Transport{ // goanalyze:ignore restricted_literal
			DialContext: d.dial,
		},
	}
}
