package transcript

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) EnrichSession(identifier string) {
	path := s.transcriptPath(identifier)
	state := s.cache.GetOrCreate(identifier)
	calls, e := tracker.Read(path, state, constant.FollowedTools)

	if e != nil {
		s.reporter.CaptureException(e)

		if state.Offset == 0 {
			return
		}
	}

	s.store.SaveTrackerState(identifier, state)
	s.RefreshSession(identifier, state)
	s.recordContextLoads(identifier, calls)
	s.index.Append(identifier, path)
}
