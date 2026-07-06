package space

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type Space struct {
	Identifier string
	Name       string
	Raw        *response.Space
}
