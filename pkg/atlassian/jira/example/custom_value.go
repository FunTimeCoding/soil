package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func CustomValue() {
	i := environment.Required(constant.JiraTestIssueEnvironment)
	f := environment.Required(constant.JiraTestFieldEnvironment)
	fmt.Printf(
		"Field value: %s\n",
		common.Jira().SetVerbose(true).MustIssue(i).CustomValue(f),
	)
}
