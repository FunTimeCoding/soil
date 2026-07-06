package service

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) PopulateCache() {
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
		state := s.cache.GetOrCreate(identifier)
		path := filepath.Join(s.harbor, entry.Name())

		if tracker.Read(path, state) != nil {
			continue
		}
	}
}
