package go_mod

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Tidy() {
	system.Run(constant.Go, constant.Mod, constant.Tidy)
}
