package stall_client

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Builder struct {
	insecure bool
	long     bool
	timeout  time.Duration
	verbose  bool
	basic    bool
	user     string
	password string
	bearer   bool
	token    string
	sendPing time.Duration
	ping     time.Duration
	observer prometheus.ObserverVec
}
