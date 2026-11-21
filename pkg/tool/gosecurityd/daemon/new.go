package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/overview"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/remediation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/tracker_source"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"github.com/funtimecoding/soil/pkg/worker"
	"path/filepath"
)

func New(
	o *option.Security,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	m := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(m)
	n := o.NationalDatabase()
	ng := peer.New(
		constant.NationalDatabasePeer,
		vulnerabilityConstant.NationalDatabaseHost,
		n,
		l,
		r,
	)
	s := o.SecurityTracker()
	sg := peer.New(
		constant.SecurityTrackerPeer,
		vulnerabilityConstant.DebianHost,
		s,
		l,
		r,
	)
	x := o.ExtendedTracker()
	xg := peer.New(
		constant.ExtendedTrackerPeer,
		vulnerabilityConstant.FreexianHost,
		x,
		l,
		r,
	)
	v := vulnerability.New(n, ng)
	d := debian.New(
		tracker_source.New(
			s,
			sg,
			o.Clock,
			filepath.Join(o.CacheDirectory, constant.DebianCache),
			true,
			constant.SupportedReleases...,
		),
		tracker_source.New(
			x,
			xg,
			o.Clock,
			filepath.Join(o.CacheDirectory, constant.ExtendedCache),
			false,
			constant.ExtendedReleases...,
		),
	)
	st := store.New(m)
	f := flag.New(st, v)
	c := score.New(st, n, ng, o.Clock)
	y := system.New(st, d, c, o.Clock, o.StaleAfter, o.TrackerStaleAfter)
	w := remediation.New(y, f)

	return &Daemon{
		database:         m,
		nationalDatabase: n,
		nationalGuard:    ng,
		securityTracker:  s,
		trackerGuard:     sg,
		extendedTracker:  x,
		extendedGuard:    xg,
		vulnerability:    v,
		debian:           d,
		flag:             f,
		system:           y,
		score:            c,
		remediation:      w,
		web:              web.New(v, d, f, overview.New(f, d), y, c, w, r),
		clock:            worker.New(d.Refresh, l, r, constant.RefreshInterval),
		scoreClock:       worker.New(c.Sync, l, r, constant.ScoreInterval),
	}
}
