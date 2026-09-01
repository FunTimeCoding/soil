package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Watch() {
	for _, i := range common.Jira().MustWatchedIssues() {
		console.Line(i.Format(constant.ColorFormat))
	}
}
