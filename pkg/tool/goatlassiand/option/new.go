package option

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
)

func New() *Atlassian {
	return &Atlassian{
		Jira: func() face.JiraClient { return jira.NewEnvironment() },
		Confluence: func() face.ConfluenceClient {
			return confluence.NewEnvironment()
		},
	}
}
