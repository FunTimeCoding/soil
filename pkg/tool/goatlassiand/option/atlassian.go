package option

import (
	confluence "github.com/funtimecoding/soil/pkg/atlassian/confluence/face"
	jira "github.com/funtimecoding/soil/pkg/atlassian/jira/face"
)

type Atlassian struct {
	Address       string
	ServiceTokens []string
	Project       []string
	Jira          func() jira.Client
	Confluence    func() confluence.Client
}
