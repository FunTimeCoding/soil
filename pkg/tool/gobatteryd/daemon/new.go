package daemon

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/target"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/worker"
	"net"
	"strconv"
)

func New(
	o *option.Exporter,
	l *logger.Logger,
	r face.Reporter,
	m *metric.Metric,
) *Daemon {
	var targets []*target.Target

	for _, i := range o.Inventory.Instances {
		c := o.Source(&i)
		targets = append(
			targets,
			target.New(
				i.Name,
				c,
				peer.New(
					fmt.Sprintf(constant.PeerFormat, i.Name),
					net.JoinHostPort(i.Host, strconv.Itoa(i.ServerPort())),
					c,
					l,
					r,
				),
			),
		)
	}

	return &Daemon{
		targets: targets,
		clock: worker.New(
			service.New(targets).Read,
			constant.PollInterval,
			l,
			m.Registry(),
			r,
		),
	}
}
