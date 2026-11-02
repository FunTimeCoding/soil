package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/query"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func collect() []*issue.Issue {
	c := common.Jira()
	defer c.Close()

	return query.Open(
		c,
		environment.Required(constant.JiraDefaultProjectNameEnvironment),
	)
}
