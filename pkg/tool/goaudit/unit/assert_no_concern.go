package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func assertNoConcern(
	t *testing.T,
	s *scan.Service,
	key string,
) {
	t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			t.Errorf("unexpected concern with key %q found", key)
		}
	}
}
