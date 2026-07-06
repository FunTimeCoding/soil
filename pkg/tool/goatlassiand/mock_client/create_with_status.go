package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) createWithStatus(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
	status string,
) (*page.Page, error) {
	identifier := fmt.Sprintf("%d", c.nextID)
	c.nextID++
	r := response.NewPage()
	r.Identifier = identifier
	r.SpaceIdentifier = spaceIdentifier
	r.ParentIdentifier = parentIdentifier
	r.Title = title
	r.Status = status
	r.Version = response.Version{Number: 1}
	r.Body = response.Body{
		Storage: response.Storage{
			Representation: constant.StorageFormat,
			Value:          page.ToStorage(markdown),
		},
	}
	r.Links = response.Links{
		WebUI: fmt.Sprintf(
			"/spaces/ops/pages/%s/%s",
			identifier,
			title,
		),
	}
	c.pages[identifier] = &entry{page: r}

	return toPage(r), nil
}
