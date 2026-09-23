package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
)

func (o *Tester) Target(identifier string) *connector.Target {
	o.t.Helper()
	targets, e := o.Coordinator.RecentSessions(50)
	assert.FatalOnError(o.t, e)

	for _, v := range targets {
		if v.Identifier == identifier {
			return v
		}
	}

	o.t.Fatalf("session %s not reported by the coordinator", identifier)

	return nil
}
