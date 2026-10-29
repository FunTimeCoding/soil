package option

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	confluenceFace "github.com/funtimecoding/soil/pkg/atlassian/confluence/face"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	jiraFace "github.com/funtimecoding/soil/pkg/atlassian/jira/face"
)

func New() *Atlassian {
	return &Atlassian{
		Jira: func() jiraFace.Client { return jira.NewEnvironment() },
		Confluence: func() confluenceFace.Client {
			return confluence.NewEnvironment()
		},
	}
}
