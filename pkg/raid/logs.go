package raid

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"time"
)

func (c *Client) Logs(
	offset *int,
	limit *int,
	start *time.Time,
	end *time.Time,
) string {
	result, e := c.client.GetLogs(
		c.context,
		&client.GetLogsParams{
			Offset: offset,
			Limit:  limit,
			Start:  start,
			End:    end,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
