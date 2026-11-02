package roster

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepEmptySessions() []error {
	closed, e := s.store.ClosedEmptyCandidates()

	if e != nil {
		return []error{e}
	}

	stale, f := s.store.StaleEmptyCandidates(s.clock().Add(-24 * time.Hour))

	if f != nil {
		return []error{f}
	}

	var result []error
	swept := 0

	for _, i := range append(closed, stale...) {
		refused, g := s.session.EmptyRefusal(&i)

		if g != nil {
			result = append(result, g)

			continue
		}

		if refused != nil {
			continue
		}

		if _, h := s.session.DeleteSession(i.Identifier, ""); h != nil {
			s.logger.Structured(
				"sweep_empty_refused",
				constant.Identifier,
				i.Identifier,
				"reason",
				h.Error(),
			)

			continue
		}

		swept++
		s.logger.Structured(
			"swept_empty_session",
			constant.Identifier,
			i.Identifier,
			constant.SessionName,
			i.Name,
			"reason",
			i.ClosedReason,
		)
	}

	if swept > 0 {
		s.notify()
	}

	return result
}
