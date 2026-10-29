package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"time"
)

func (c *Client) Resolve(
	identifier int,
	end time.Time,
) {
	for i := range c.alerts {
		if c.alerts[i].Identifier == identifier {
			c.alerts[i].End = new(end.Format(constant.DateFormat))
			c.alerts[i].Status = client.Resolved
		}
	}
}
