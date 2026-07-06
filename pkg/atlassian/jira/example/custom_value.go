package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func CustomValue() {
	i := environment.Required(constant.TestIssueEnvironment)
	f := environment.Required(constant.TestFieldEnvironment)
	fmt.Printf(
		"Field value: %s\n",
		common.Jira().SetVerbose(true).MustIssue(i).CustomValue(f),
	)
}
