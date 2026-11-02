package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Silence() {
	f := consoleConstant.ColorFormat.Copy().Tag(consoleConstant.TagState)
	c := common.Alertmanager()
	defer c.Close()

	for _, a := range c.MustSilences(true) {
		console.Line(a.Format(f))
	}
}
