package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Watch() {
	c := common.Jira()
	defer c.Close()

	for _, i := range c.MustWatchedIssues() {
		console.Line(i.Format(consoleConstant.ColorFormat))
	}
}
