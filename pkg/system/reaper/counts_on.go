package reaper

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
)

func CountsOn(y *prometheus.Registry) bool {
	var registered prometheus.AlreadyRegisteredError

	return errors.As(y.Register(NewCounter()), &registered)
}
