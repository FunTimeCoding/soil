package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func Concern(
	t *testing.T,
	key string,
	s *scan.Service,
) {
	t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			return
		}
	}

	t.Errorf("expected concern with key %q not found", key)
}
