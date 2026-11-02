package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Terraform,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d, constant.RunsTable)
	s := store.New(d, constant.RunsTable)
	s.AbandonRunning()
	result := &Service{
		database:   d,
		store:      s,
		reaper:     reaper.New(r),
		downstream: o.DownstreamClients(),
	}
	var targets []provisionFace.Downstream

	for _, c := range result.downstream {
		g := peer.New(
			join.Colon(constant.DownstreamPeer, c.Address()),
			c.Address(),
			c,
			l,
			r,
		)
		result.guards = append(result.guards, g)
		targets = append(targets, guarded(c, g))
	}

	result.runner = runner.New(
		o,
		s,
		l,
		r,
		result.reaper,
		y,
		o.Cluster(),
		targets,
	)

	return result
}
