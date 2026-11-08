package roster

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) RecordCompaction(
	sessionIdentifier string,
	phase constant.CompactionPhase,
	trigger string,
) error {
	d, e := compactionDelivery(phase)

	if e != nil {
		return e
	}

	r, found, f := s.store.FindSession(sessionIdentifier)

	if f != nil {
		return f
	}

	if !found {
		return nil
	}

	if g := s.store.LogEvent(
		sessionIdentifier,
		constant.Compaction,
		r.CallsignValue(),
		map[string]string{
			constant.Phase:   string(phase),
			constant.Trigger: trigger,
		},
	); g != nil {
		return g
	}
	s.notify()

	return s.tellWatchers(
		sessionIdentifier,
		compactionNotice(r.CallsignValue(), phase, trigger, r.Topic),
		d,
	)
}
