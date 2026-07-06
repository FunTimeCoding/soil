package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepCallsignRelease() {
	cutoff := s.clock().Add(-72 * time.Hour)
	sessions := s.store.SweepCallsignRelease(cutoff)

	for _, e := range sessions {
		s.logger.Structured(
			"callsign_released",
			constant.Callsign,
			e.CallsignValue(),
			constant.Identifier,
			e.Identifier,
		)
	}
}
