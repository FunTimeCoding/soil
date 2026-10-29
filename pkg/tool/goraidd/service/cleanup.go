package service

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"time"
)

func (s *Service) Cleanup() {
	cutoff := time.Now().AddDate(0, 0, -14)
	deleted := 0

	for _, n := range system.ReadDirectory(s.elitePath) {
		i, e := n.Info()

		if e != nil || i.ModTime().After(cutoff) {
			continue
		}

		system.RemoveFile(filepath.Join(s.elitePath, n.Name()))
		deleted++
	}

	if deleted > 0 {
		s.logger.Structured("cleanup_completed", "deleted", deleted)
	}
}
