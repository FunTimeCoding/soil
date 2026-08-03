package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Watch() {
	for _, i := range common.Jira().MustWatchedIssues() {
		fmt.Println(i.Format(constant.ColorFormat))
	}
}
