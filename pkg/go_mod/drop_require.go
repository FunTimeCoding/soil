package go_mod

import (
	"fmt"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func dropRequire(mod string) {
	system.Run(
		soil.Go,
		soil.Mod,
		soil.Edit,
		fmt.Sprintf("-droprequire=%s", mod),
	)
}
