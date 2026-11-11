package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/option"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/service"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/worker"
)

func New(
	o *option.Patrol,
	l *logger.Logger,
	r face.Reporter,
	m *metric.Metric,
) *Daemon {
	var guards []*peer.Guard
	var targets []*target.Target

	for _, host := range o.Hosts {
		c := o.Source(host)
		g := peer.New(join.Space(constant.PatrolPeer, host), host, c, l, r)
		guards = append(guards, g)
		targets = append(targets, target.New(host, c, g))
	}

	s := service.New(targets...)

	return &Daemon{
		guards:  guards,
		service: s,
		clock:   worker.New(s.Round, constant.PollInterval, l, m.Registry(), r),
	}
}
