package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Watch() {
	for _, i := range common.Jira().MustWatchedIssues() {
		fmt.Println(i.Format(option.Color))
	}
}
