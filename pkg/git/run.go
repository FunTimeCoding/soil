package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Run(s ...string) {
	system.Run(append([]string{constant.Command}, s...)...)
}
