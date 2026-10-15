package observe_transport

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type Transport struct {
	observer prometheus.ObserverVec
	base     http.RoundTripper
}
