package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
)

func New() *Log {
	return &Log{
		Retention: constant.Retention,
		Alertmanager: func() face.Client {
			return alertmanager.NewEnvironment()
		},
	}
}
