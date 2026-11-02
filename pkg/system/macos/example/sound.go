package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	soundConstant "github.com/funtimecoding/soil/pkg/sound/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Sound() {
	console.Line(system.Run(soundConstant.Afplay, soundConstant.SosumiPath))
	console.Line(system.Run(soundConstant.Afplay, soundConstant.TinkPath))
}
