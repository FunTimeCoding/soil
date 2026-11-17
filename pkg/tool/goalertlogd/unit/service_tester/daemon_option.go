package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
)

func (o *Tester) daemonOption() *option.Log {
	result := option.New()
	result.LitePath = o.Path
	result.Retention = o.retention
	result.Clock = o.clock
	result.Alertmanager = func() face.AlertmanagerClient { return o.Alertmanager }

	return result
}
