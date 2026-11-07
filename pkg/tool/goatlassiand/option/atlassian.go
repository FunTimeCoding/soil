package option

import (
	"github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Atlassian struct {
	option.Web
	Project    []string
	Jira       func() face.JiraClient
	Confluence func() face.ConfluenceClient
}
