package project

import "gitlab.com/gitlab-org/api/client-go/v3"

func New(v *gitlab.Project) *Project {
	return &Project{
		Identifier: v.ID,
		Namespace:  v.Namespace.Path,
		Name:       v.Path,
		Link:       v.WebURL,
		Raw:        v,
	}
}
