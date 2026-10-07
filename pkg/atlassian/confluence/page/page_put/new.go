package page_put

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/types/page_version"
)

func New(
	identifier string,
	title string,
	body string,
	version int,
	message string,
) *Put {
	return &Put{
		Identifier: identifier,
		Status:     "current",
		Title:      title,
		Body: response.Storage{
			Representation: constant.ConfluenceStorageFormat,
			Value:          body,
		},
		Version: page_version.Version{Number: version, Message: message},
	}
}
