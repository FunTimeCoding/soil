package search

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Search() {
	p := environment.Required(constant.DefaultProjectNameEnvironment)
	j := common.Jira()
	f := constant.Format
	searchAndy(j, p, f)
	searchOwn(j, p)
	searchOwnFull(j, p, f)
}
