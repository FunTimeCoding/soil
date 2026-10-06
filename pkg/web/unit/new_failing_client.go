package unit

import (
	"context"
	"net"
	"net/http"
	"sync/atomic"
)

func newFailingClient(failure error) (*http.Client, *atomic.Int32) {
	var attempts atomic.Int32

	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(
				context.Context,
				string,
				string,
			) (net.Conn, error) {
				attempts.Add(1)

				return nil, failure
			},
		},
	}, &attempts
}
