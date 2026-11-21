package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func (o *Tester) LabeledFleet() *Tester {
	o.t.Helper()
	o.SecurityTracker.SeedKnown()
	assert.Count(o.t, 0, o.Debian().Refresh(o.t.Context()))

	for name, labels := range map[string]map[string]string{
		"alfa":    {"exposure": "internet", "tenant": "one"},
		"bravo":   {"exposure": "internet", "tenant": "two"},
		"charlie": {"exposure": "internal", "tenant": "one"},
		"delta":   {"exposure": "internal"},
	} {
		v := system.New(name, "trixie", 0)
		v.Labels = labels
		assert.FatalOnError(
			o.t,
			o.System().Put(
				v,
				[]installed.Installed{
					*installed.New(
						"libtango5",
						"8.4.0-1",
						"tango-utils",
						"8.4.0-1",
						"8.6.1-1",
					),
				},
			),
		)
	}

	return o
}
