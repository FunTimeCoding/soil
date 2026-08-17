package opnsense

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/log_entry"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"strconv"
)

func (c *Client) Log(limit int) ([]*log_entry.Entry, error) {
	b, e := c.basic.Get(
		constant.LogRead,
		map[string]string{"limit": strconv.Itoa(limit)},
	)

	if e != nil {
		return nil, e
	}

	var out []response.LogEntry

	if f := json.Unmarshal(b, &out); f != nil {
		return nil, f
	}

	return log_entry.NewSlice(out), nil
}
