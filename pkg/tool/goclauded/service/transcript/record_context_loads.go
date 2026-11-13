package transcript

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"

func (s *Service) recordContextLoads(
	identifier string,
	calls []*tool_call.Call,
) {
	loads, profiles := classifyContextLoads(identifier, calls)

	for _, readable := range profiles {
		s.noteProfile(readable)
	}

	s.notePersistedProfiles(calls)

	if len(loads) == 0 {
		return
	}

	if e := s.store.SaveContextLoads(loads); e != nil {
		s.reporter.CaptureException(e)
	}
}
