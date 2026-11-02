package service

import (
	"os"
	"path/filepath"
	"time"
)

func (s *Service) Cleanup() []error {
	entries, e := os.ReadDir(s.elitePath)

	if e != nil {
		return []error{e}
	}

	cutoff := time.Now().AddDate(0, 0, -14)
	deleted := 0
	var result []error

	for _, n := range entries {
		i, f := n.Info()

		if f != nil || i.ModTime().After(cutoff) {
			continue
		}

		if g := os.Remove(filepath.Join(s.elitePath, n.Name())); g != nil {
			result = append(result, g)

			continue
		}

		deleted++
	}

	if deleted > 0 {
		s.logger.Structured("cleanup_completed", "deleted", deleted)
	}

	return result
}
