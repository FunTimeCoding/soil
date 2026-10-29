package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"time"
)

type Log struct {
	Address         string
	ServiceTokens   []string
	MetricAddress   string
	PostgresLocator string
	LitePath        string
	Retention       time.Duration
	Alertmanager    func() face.Client
}
