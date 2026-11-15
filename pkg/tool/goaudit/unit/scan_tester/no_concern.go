package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func NoConcern(
	t *testing.T,
	key string,
	s *scan.Service,
) {
	t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			t.Errorf("unexpected concern with key %q found", key)
		}
	}
}
