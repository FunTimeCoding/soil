package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/ceph/constant"
	"github.com/funtimecoding/soil/pkg/ceph/goc"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	if false {
		// Only works for the running process
		environment.Set(constant.ConfigurationEnvironment, upper.Alfa)
		environment.Set(constant.ArgumentEnvironment, upper.Bravo)
	}

	if false {
		// Not sure if this works
		goc.SetEnvironmentEscape(constant.ConfigurationEnvironment, upper.Alfa)
		goc.SetEnvironmentEscape(constant.ArgumentEnvironment, upper.Bravo)
	}

	if false {
		// Not working
		fmt.Printf(
			"Get escape: %s\n",
			goc.GetEnvironmentEscape(constant.ConfigurationEnvironment),
		)
		fmt.Printf(
			"Get escape: %s\n",
			goc.GetEnvironmentEscape(constant.ArgumentEnvironment),
		)
	}
}
