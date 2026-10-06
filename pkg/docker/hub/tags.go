package hub

import (
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/docker/hub/tag"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Tags(image string) ([]*tag.Tag, error) {
	var result tag.ListResponse

	if e := c.requester.Notation(
		request.Get(join.Empty(image, "/tags")).WithParameter(
			constant.PageSizeParameter,
			constant.PageSize,
		),
		&result,
	); e != nil {
		return nil, e
	}

	return tag.NewSlice(result.Results), nil
}
