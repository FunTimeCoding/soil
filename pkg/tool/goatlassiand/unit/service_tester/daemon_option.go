package service_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
)

func (o *Tester) daemonOption() *option.Atlassian {
	result := option.New()
	result.Project = o.project
	result.Jira = func() face.JiraClient { return o.Jira }
	result.Confluence = func() face.ConfluenceClient { return o.Confluence }

	return result
}
