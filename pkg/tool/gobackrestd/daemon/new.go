package daemon

import (
	backrestConstant "github.com/funtimecoding/soil/pkg/backrest/constant"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/service"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/worker"
)

func New(
	o *option.Exporter,
	l *logger.Logger,
	r face.Reporter,
	m *metric.Metric,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.BackrestPeer, backrestConstant.Executable, c, l, r)
	s := service.New(c, g)

	return &Daemon{
		guard:   g,
		service: s,
		clock: worker.New(
			s.Latest,
			constant.PollInterval,
			l,
			m.Registry(),
			r,
		),
	}
}
