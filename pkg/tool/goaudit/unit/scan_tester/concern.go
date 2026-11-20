package scan_tester

import "github.com/funtimecoding/soil/pkg/tool/goaudit/scan"

func (o *Tester) Concern(
	key string,
	s *scan.Service,
) {
	o.t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			return
		}
	}

	o.t.Errorf("expected concern with key %q not found", key)
}
