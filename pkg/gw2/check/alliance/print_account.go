package alliance

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func PrintAccount() {
	console.Format(
		"%+v",
		gw2.New(environment.Required("GW2_TEST_TOKEN")).Account(),
	)
}
