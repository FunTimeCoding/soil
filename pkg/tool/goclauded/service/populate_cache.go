package service

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) PopulateCache() {
	for identifier, state := range s.store.TrackerStates() {
		s.cache.Set(identifier, state)
	}

	entries, e := os.ReadDir(s.harbor)

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		) {
			continue
		}

		identifier := strings.TrimSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		)
		i, f := entry.Info()

		if f != nil {
			continue
		}

		state := s.cache.GetOrCreate(identifier)

		if state.Offset == i.Size() {
			continue
		}

		path := filepath.Join(s.harbor, entry.Name())

		if tracker.Read(path, state) != nil {
			continue
		}

		s.store.SaveTrackerState(identifier, state)
	}
}
