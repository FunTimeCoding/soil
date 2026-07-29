package search

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Search() {
	p := environment.Required(constant.JiraDefaultProjectNameEnvironment)
	j := common.Jira()
	f := constant.JiraFormat
	searchAndy(j, p, f)
	searchOwn(j, p)
	searchOwnFull(j, p, f)
}
