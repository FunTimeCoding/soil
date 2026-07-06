package hub

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/docker/hub/constant"
	"github.com/funtimecoding/soil/pkg/docker/hub/tag"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Client) Tags(image string) []*tag.Tag {
	r := web.NewGet(
		c.base.Copy().Path(
			join.Empty(image, "/tags"),
		).Set(constant.PageSizeParameter, constant.PageSize).String(),
	)
	r.Header.Set(webConstant.Accept, webConstant.Object)
	response := web.Send(web.Client(), r)
	defer errors.PanicClose(response.Body)
	var result tag.ListResponse
	errors.PanicOnError(json.NewDecoder(response.Body).Decode(&result))

	return tag.NewSlice(result.Results)
}
