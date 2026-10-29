package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service/outcome"
	"time"
)

func (s *Service) sightLeases(
	x context.Context,
	g *gazetteer.Gazetteer,
) *outcome.Outcome {
	started := time.Now()
	v, e := s.lease.Collect(x, g)

	if e != nil {
		return outcome.Failed(constant.SourceLease, time.Since(started), e)
	}

	swept, f := s.saveSightings(constant.SourceLease, v)

	if f != nil {
		return outcome.Failed(constant.SourceLease, time.Since(started), f)
	}

	return outcome.Sighted(
		constant.SourceLease,
		len(v),
		swept,
		time.Since(started),
	)
}
