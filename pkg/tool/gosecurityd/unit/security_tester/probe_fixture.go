package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
)

func ProbeFixture(name string) string {
	return fixture.Read(
		systemConstant.VulnerabilityPath,
		join.Empty("probe-", name, ".txt"),
	)
}
