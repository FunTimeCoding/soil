package alliance

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func PrintAccount() {
	fmt.Printf(
		"%+v",
		gw2.New(environment.Required("GW2_TEST_TOKEN")).Account(),
	)
}
