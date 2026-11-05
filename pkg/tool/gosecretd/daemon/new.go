package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
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
	c := o.Source()
	g := peer.New(constant.OpenbaoPeer, c.Address(), c, l, r)
	v := service.New(c, g)

	return &Daemon{
		client:  c,
		guard:   g,
		service: v,
		watch:   watch.New(v.Status, y, l, r),
	}
}
