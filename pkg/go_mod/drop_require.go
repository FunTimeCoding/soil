package go_mod

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func dropRequire(mod string) {
	system.Run(
		constant.Go,
		constant.Mod,
		constant.Edit,
		fmt.Sprintf("-droprequire=%s", mod),
	)
}
