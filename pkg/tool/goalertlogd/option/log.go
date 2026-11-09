package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Log struct {
	option.Web
	Metric          option.Web
	PostgresLocator string
	LitePath        string
	Retention       time.Duration
	Clock           func() time.Time
	Alertmanager    func() face.AlertmanagerClient
}
