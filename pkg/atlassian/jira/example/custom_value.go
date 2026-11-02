package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func CustomValue() {
	i := environment.Required(constant.JiraTestIssueEnvironment)
	f := environment.Required(constant.JiraTestFieldEnvironment)
	c := common.Jira().SetVerbose(true)
	defer c.Close()
	console.Format("Field value: %s\n", c.MustIssue(i).CustomValue(f))
}
