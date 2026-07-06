package issue

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/query"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func collect() []*issue.Issue {
	return query.Open(
		common.Jira(),
		environment.Required(constant.DefaultProjectNameEnvironment),
	)
}
