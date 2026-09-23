package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func assertConcern(
	t *testing.T,
	s *scan.Service,
	key string,
) {
	t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			return
		}
	}

	t.Errorf("expected concern with key %q not found", key)
}
