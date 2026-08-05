package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) RecordContext(
	identifier string,
	percent int,
	window int,
	model string,
) {
	updates := map[string]any{
		constant.ContextPercentColumn: percent,
	}

	if window > 0 {
		updates[constant.ContextWindowColumn] = window
	}

	if model != "" {
		updates[constant.ModelColumn] = model
	}

	s.store.UpdateFields(identifier, updates)
}
