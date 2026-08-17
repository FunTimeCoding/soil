package service

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/soil/pkg/strings/join"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"path/filepath"
)

func (s *Service) EnrichSession(identifier string) {
	path := filepath.Join(
		s.harbor,
		join.Empty(identifier, constant.NotationLogExtension),
	)
	state := s.cache.GetOrCreate(identifier)
	calls, e := tracker.Read(path, state, goclaudedConstant.FollowedTools)

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	s.store.SaveTrackerState(identifier, state)
	s.RefreshSession(identifier, state)
	s.recordContextLoads(identifier, calls)
}
