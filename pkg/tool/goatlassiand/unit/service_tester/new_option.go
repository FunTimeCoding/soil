package service_tester

import (
	confluence "github.com/funtimecoding/soil/pkg/atlassian/confluence/face"
	jira "github.com/funtimecoding/soil/pkg/atlassian/jira/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
)

func NewOption(
	j jira.Client,
	c confluence.Client,
	project []string,
) *option.Atlassian {
	result := option.New()
	result.Project = project
	result.Jira = func() jira.Client { return j }
	result.Confluence = func() confluence.Client { return c }

	return result
}
