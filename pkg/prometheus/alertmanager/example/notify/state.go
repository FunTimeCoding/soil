package notify

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

type State struct {
	Loaded bool
	Alerts []*alert.Alert
}
