package sweeper

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest"
	"time"
)

func (s *Sweeper) Sweep() {
	pruned, e := s.store.Prune(time.Now().Add(-s.window))

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	if len(pruned) == 0 {
		return
	}

	dropped := map[string][]string{}

	for _, v := range pruned {
		dropped[v.Callsign] = append(dropped[v.Callsign], v.Label())
		s.release(v)
	}

	for callsign, label := range dropped {
		s.notifier.Notify(callsign, digest.Dropped(label, s.window))
	}

	s.logger.Structured("subscription_purge", "count", len(pruned))
}
