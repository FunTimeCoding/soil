package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/parse"
)

func (s *Service) SampleBluetooth() []error {
	output, e := s.command(
		constant.SystemProfiler,
		constant.BluetoothSection,
		constant.ProfilerNotation,
	)

	if e != nil {
		s.logger.Structured("system profiler failed", "error", e.Error())

		return nil
	}

	if f := s.recordChanges(
		constant.BluetoothKind,
		parse.Bluetooth(output),
	); f != nil {
		return []error{f}
	}

	return nil
}
