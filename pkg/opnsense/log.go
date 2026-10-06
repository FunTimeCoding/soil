package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"strconv"
)

func (c *Client) Log(limit int) ([]*log_entry.Entry, error) {
	var out []response.LogEntry

	if e := c.basic.Get(
		constant.LogRead,
		map[string]string{"limit": strconv.Itoa(limit)},
		&out,
	); e != nil {
		return nil, e
	}

	return log_entry.NewSlice(out), nil
}
