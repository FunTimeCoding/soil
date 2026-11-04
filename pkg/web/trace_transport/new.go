package trace_transport

import "net/http"

func New(t *http.Transport) *Transport {
	return &Transport{transport: t}
}
