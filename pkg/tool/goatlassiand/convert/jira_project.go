package convert

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func JiraProject(p *jira.Project) *server.JiraProject {
	result := &server.JiraProject{Key: p.Key, Name: p.Name}

	if p.Description != "" {
		result.Description = &p.Description
	}

	if p.ProjectCategory.Name != "" {
		result.Type = &p.ProjectCategory.Name
	}

	if p.Self != "" {
		result.Link = &p.Self
	}

	return result
}
