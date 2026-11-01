package service

import (
	"context"
	"time"
)

func (s *Service) saveImpression(
	x context.Context,
	content string,
	source string,
	capturedAt time.Time,
) error {
	return s.memoryGuard.Attempt(
		x,
		func(y context.Context) error {
			return s.memory.PostImpressions(y, content, source, capturedAt)
		},
	)
}
