package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) reportRefusedImpression(
	e error,
	content string,
	source string,
	capturedAt time.Time,
) {
	s.reporter.CaptureWithContext(
		e,
		constant.ImpressionContext,
		map[string]any{
			constant.ImpressionSource:     source,
			constant.ImpressionContent:    content,
			constant.ImpressionCapturedAt: capturedAt.UTC().Format(time.RFC3339),
		},
	)
}
