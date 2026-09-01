package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/sound/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Sound() {
	console.Line(system.Run(constant.Afplay, constant.SosumiPath))
	console.Line(system.Run(constant.Afplay, constant.TinkPath))
}
