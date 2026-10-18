package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"time"
)

func (s *Service) ForwardImpression(
	x context.Context,
	content string,
	source string,
	capturedAt time.Time,
) {
	e := s.saveImpression(x, content, source, capturedAt)

	if e == nil {
		return
	}

	if !unavailable.Is(e) {
		s.reportRefusedImpression(e, content, source, capturedAt)

		return
	}

	if f := s.store.QueueImpression(source, content, capturedAt); f != nil {
		s.reporter.CaptureException(f)
	}
}
