package observe_transport

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

func New(
	o prometheus.ObserverVec,
	base http.RoundTripper,
) *Transport {
	return &Transport{observer: o, base: base}
}
