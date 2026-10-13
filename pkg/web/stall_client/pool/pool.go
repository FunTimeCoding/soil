package pool

import "net/http"

type Pool struct {
	transport *http.Transport
	client    *http.Client
	halting   *http.Client
}
