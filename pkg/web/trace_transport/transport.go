package trace_transport

import "net/http"

type Transport struct {
	transport *http.Transport
}
