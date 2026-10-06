package raid_parser

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"time"
)

func (c *Client) Generate(
	files []string,
	date *time.Time,
) (string, error) {
	result, e := c.client.PostGenerate(
		c.context,
		client.PostGenerateJSONRequestBody{Files: files, Date: date},
	)

	if e != nil {
		return "", e
	}

	body := web.ReadString(result)

	if !web.ResponseOkay(result) {
		return "", unexpected.Format(
			"raid parser generate status: %d: %s",
			result.StatusCode,
			body,
		)
	}

	return body, nil
}
