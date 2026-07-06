package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/logs"
)

func (c *Client) Logs(marker string) []logs.Log {
	r, e := c.userClient.Log.ListLogFiles(
		c.context,
		&logs.ListLogFilesRequest{Marker: marker},
	)
	errors.PanicOnError(e)

	return r.Logs
}
