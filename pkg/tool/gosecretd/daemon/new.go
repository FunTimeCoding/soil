package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Secret,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	s := o.Source()
	v := service.New(s)

	return &Daemon{
		source:  s,
		service: v,
		watch:   watch.New(v.Status, y, l, r),
	}
}
