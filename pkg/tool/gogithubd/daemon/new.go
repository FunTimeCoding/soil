package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
)

func New(
	o *option.Exporter,
	l *logger.Logger,
	r face.Reporter,
	m *metric.Metric,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.GitHubPeer, constant.GitHubAddress, c, l, r)
	s := service.New(c, g, o.Owner)

	return &Daemon{
		client:  c,
		guard:   g,
		service: s,
		clock: worker.New(
			s.Latest,
			o.Owner,
			constant.PollInterval,
			l,
			m.Registry(),
			r,
		),
	}
}
