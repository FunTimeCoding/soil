package option

import "github.com/funtimecoding/soil/pkg/atlassian/face"

type Atlassian struct {
	Address       string
	ServiceTokens []string
	Project       []string
	Jira          func() face.JiraClient
	Confluence    func() face.ConfluenceClient
}
