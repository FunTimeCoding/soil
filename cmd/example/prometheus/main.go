package main

import (
	grafana "github.com/funtimecoding/soil/pkg/grafana/example"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/example"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/example/notify"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/example"
	loki "github.com/funtimecoding/soil/pkg/prometheus/loki/example"
)

func main() {
	loki.QueryRange()

	if false {
		example.Alert()
		example.Create()
		example.DeleteSilence()
		notify.Notify()
		example.SetSilence()
		example.Silence()
		example.Status()
		grafana.Read()
		loki.Label()
		loki.Query()
		loki.Series()
		loki.Statistic()
		loki.Write()
		prometheus.Label()
		prometheus.LabelName()
		prometheus.Meta()
		prometheus.Metric()
		prometheus.Query()
		prometheus.Rule()
		prometheus.Series()
		prometheus.Status()
		prometheus.Target()
	}
}
