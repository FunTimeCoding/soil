package page_post

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

type Post struct {
	SpaceIdentifier  string           `json:"spaceId"`
	Status           string           `json:"status"`
	Title            string           `json:"title"`
	ParentIdentifier string           `json:"parentId"`
	Body             response.Storage `json:"body"`
}
