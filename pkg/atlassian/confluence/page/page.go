package page

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type Page struct {
	Identifier string
	Name       string
	Link       string
	TinyLink   string
	Raw        *response.Page
}
