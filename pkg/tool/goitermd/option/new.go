package option

import (
	"github.com/funtimecoding/soil/pkg/iterm"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
)

func New() *Iterm {
	return &Iterm{
		Source:      func() face.ItermSource { return iterm.NewEnvironment() },
		KeyInterval: constant.KeyInterval,
	}
}
