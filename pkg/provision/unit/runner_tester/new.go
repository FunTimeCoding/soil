package runner_tester

import "testing"

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	s := &setup{hold: func() {}}

	for _, f := range o {
		f(s)
	}

	return build(t, s.hold, s.registry, s.downstream, s.copyModes)
}
