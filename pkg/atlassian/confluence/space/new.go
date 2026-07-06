package space

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func New(p *response.Space) *Space {
	return &Space{Identifier: p.Identifier, Name: p.Name, Raw: p}
}
