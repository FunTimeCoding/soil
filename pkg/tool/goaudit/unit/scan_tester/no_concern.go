package scan_tester

import "github.com/funtimecoding/soil/pkg/tool/goaudit/scan"

func (o *Tester) NoConcern(
	key string,
	s *scan.Service,
) {
	o.t.Helper()

	for _, c := range s.Concerns {
		if c.Key == key {
			o.t.Errorf("unexpected concern with key %q found", key)
		}
	}
}
