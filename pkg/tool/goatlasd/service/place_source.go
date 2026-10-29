package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service/outcome"
	"time"
)

func (s *Service) placeSource(
	x context.Context,
	c face.Collector,
	g *gazetteer.Gazetteer,
) *outcome.Outcome {
	started := time.Now()
	v, e := c.Collect(x, g)

	if e != nil {
		return outcome.Failed(c.Source(), time.Since(started), e)
	}

	swept, f := s.savePlacements(c.Source(), v)

	if f != nil {
		return outcome.Failed(c.Source(), time.Since(started), f)
	}

	return outcome.Placed(c.Source(), len(v), swept, time.Since(started))
}
