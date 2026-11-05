package option

import (
	"github.com/funtimecoding/soil/pkg/iterm"
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
)

func New() *Iterm {
	return &Iterm{
		Source:      func() face.ItermClient { return iterm.NewEnvironment() },
		KeyInterval: constant.KeyInterval,
	}
}
