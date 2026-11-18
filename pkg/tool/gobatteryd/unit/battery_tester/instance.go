package battery_tester

import (
	"github.com/funtimecoding/soil/pkg/nut/fixture"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/inventory/instance"
)

func Instance(
	name string,
	s *fixture.Server,
) *instance.Instance {
	return instance.New(name, s.Host(), s.Port())
}
