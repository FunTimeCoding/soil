package service_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
)

func NewOption(
	j face.JiraClient,
	c face.ConfluenceClient,
	project []string,
) *option.Atlassian {
	result := option.New()
	result.Project = project
	result.Jira = func() face.JiraClient { return j }
	result.Confluence = func() face.ConfluenceClient { return c }

	return result
}
