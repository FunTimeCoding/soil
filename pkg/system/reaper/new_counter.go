package reaper

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func NewCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "reaper_reaped_total",
			Help: "Zombie children reaped, by the command they ran.",
		},
		[]string{constant.ReapedCommand},
	)
}
