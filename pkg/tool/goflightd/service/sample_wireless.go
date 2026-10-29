package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/parse"
)

func (s *Service) SampleWireless() []error {
	output, e := s.wireless()

	if e != nil {
		s.logger.Structured("wdutil failed", "error", e.Error())

		return nil
	}

	if f := s.recordChanges(
		constant.WirelessKind,
		parse.Wireless(output),
	); f != nil {
		return []error{f}
	}

	return nil
}
