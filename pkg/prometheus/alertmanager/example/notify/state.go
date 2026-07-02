package notify

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

type State struct {
	Loaded bool
	Alerts []*alert.Alert
}
