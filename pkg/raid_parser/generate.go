package raid_parser

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"time"
)

func (c *Client) Generate(
	files []string,
	date *time.Time,
) string {
	result, e := c.client.PostGenerate(
		c.context,
		client.PostGenerateJSONRequestBody{
			Files: files,
			Date:  date,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
