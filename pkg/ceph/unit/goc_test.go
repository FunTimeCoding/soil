package unit

import (
	"github.com/funtimecoding/soil/pkg/ceph/constant"
	"github.com/funtimecoding/soil/pkg/ceph/goc"
	"github.com/funtimecoding/soil/pkg/console"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	if false {
		// Only works for the running process
		environment.Set(constant.ConfigurationEnvironment, strings.UpperAlfa)
		environment.Set(constant.ArgumentEnvironment, strings.UpperBravo)
	}

	if false {
		// Not sure if this works
		goc.SetEnvironmentEscape(
			constant.ConfigurationEnvironment,
			strings.UpperAlfa,
		)
		goc.SetEnvironmentEscape(
			constant.ArgumentEnvironment,
			strings.UpperBravo,
		)
	}

	if false {
		// Not working
		console.Format(
			"Get escape: %s\n",
			goc.GetEnvironmentEscape(constant.ConfigurationEnvironment),
		)
		console.Format(
			"Get escape: %s\n",
			goc.GetEnvironmentEscape(constant.ArgumentEnvironment),
		)
	}
}
