package pool

import (
	"net/http"
	"time"
)

func New(
	t *http.Transport,
	r http.RoundTripper,
	timeout time.Duration,
) *Pool {
	return &Pool{
		transport: t,
		client:    &http.Client{Transport: r, Timeout: timeout},
		halting: &http.Client{
			Transport: r,
			Timeout:   timeout,
			CheckRedirect: func(
				*http.Request,
				[]*http.Request,
			) error {
				return http.ErrUseLastResponse
			},
		},
	}
}
