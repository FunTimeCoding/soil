package convert

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func JiraProjects(v *jira.ProjectList) []*server.JiraProject {
	result := make([]*server.JiraProject, 0, len(*v))

	for _, p := range *v {
		r := &server.JiraProject{Key: p.Key, Name: p.Name}

		if p.ProjectTypeKey != "" {
			r.Type = &p.ProjectTypeKey
		}

		if p.Self != "" {
			r.Link = &p.Self
		}

		result = append(result, r)
	}

	return result
}
