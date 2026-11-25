package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store"
	"github.com/funtimecoding/soil/pkg/watch/notifier"
	"github.com/funtimecoding/soil/pkg/watch/pending_digest"
)

func (s *Service) WithWatches(
	t *store.Store,
	c goclauded.Client,
	g *peer.Guard,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	s.store = t
	s.roster = c
	s.rosterGuard = g
	s.digests = notifier.New(c, constant.Identity.Name(), g)
	s.pending = pending_digest.New(
		constant.WatchInterval,
		constant.PendingRetryLimit,
	)
	s.outages = outage_tracker.New(r)
	s.logger = l

	return s
}
