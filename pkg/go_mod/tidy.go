package go_mod

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Tidy() {
	system.Run(soil.Go, soil.Mod, soil.Tidy)
}
