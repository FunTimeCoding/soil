package page_put

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/types/page_version"
)

type Put struct {
	Identifier string               `json:"id"`
	Status     string               `json:"status"`
	Title      string               `json:"title"`
	Body       response.Storage     `json:"body"`
	Version    page_version.Version `json:"version"`
}
