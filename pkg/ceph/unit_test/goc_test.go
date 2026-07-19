package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/ceph/goc"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	if false {
		// Only works for the running process
		environment.Set(goc.ConfigurationEnvironment, upper.Alfa)
		environment.Set(goc.ArgumentEnvironment, upper.Bravo)
	}

	if false {
		// Not sure if this works
		goc.SetEnvironmentEscape(goc.ConfigurationEnvironment, upper.Alfa)
		goc.SetEnvironmentEscape(goc.ArgumentEnvironment, upper.Bravo)
	}

	if false {
		// Not working
		fmt.Printf(
			"Get escape: %s\n",
			goc.GetEnvironmentEscape(goc.ConfigurationEnvironment),
		)
		fmt.Printf(
			"Get escape: %s\n",
			goc.GetEnvironmentEscape(goc.ArgumentEnvironment),
		)
	}
}
