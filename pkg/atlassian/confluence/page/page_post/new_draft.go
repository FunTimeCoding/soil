package page_post

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func NewDraft(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	body string,
) *Post {
	return &Post{
		SpaceIdentifier:  spaceIdentifier,
		ParentIdentifier: parentIdentifier,
		Title:            title,
		Status:           constant.ConfluenceDraftStatus,
		Body: response.Storage{
			Representation: constant.ConfluenceStorageFormat,
			Value:          body,
		},
	}
}
