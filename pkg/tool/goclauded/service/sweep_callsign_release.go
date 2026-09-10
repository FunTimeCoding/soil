package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) sweepCallsignRelease() {
	s.store.SweepCallsignRelease(s.clock().Add(-constant.CallsignReleaseWindow))
}
