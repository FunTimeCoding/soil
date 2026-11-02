package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service/peered"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func peeredBuilder(
	o *option.Alertmanager,
	l *logger.Logger,
	r face.Reporter,
) func(*instance.Instance) (*peered.Peered, error) {
	return func(i *instance.Instance) (*peered.Peered, error) {
		a, e := o.Alertmanager(i)

		if e != nil {
			return nil, e
		}

		p, f := o.Prometheus(i)

		if f != nil {
			a.Close()

			return nil, f
		}

		return peered.New(
			a,
			p,
			peer.New(
				join.Colon(constant.AlertmanagerPeer, i.Name),
				i.Alertmanager.Host,
				a,
				l,
				r,
			),
			peer.New(
				join.Colon(constant.PrometheusPeer, i.Name),
				i.Prometheus.Host,
				p,
				l,
				r,
			),
		), nil
	}
}
