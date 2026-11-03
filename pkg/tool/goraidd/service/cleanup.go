package service

import (
	"context"
	"os"
	"path/filepath"
	"time"
)

func (s *Service) Cleanup(x context.Context) []error {
	entries, e := os.ReadDir(s.elitePath)

	if e != nil {
		return []error{e}
	}

	cutoff := time.Now().AddDate(0, 0, -14)
	deleted := 0
	var result []error

	for _, n := range entries {
		if x.Err() != nil {
			break
		}

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
