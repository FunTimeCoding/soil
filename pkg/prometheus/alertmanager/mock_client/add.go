package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (c *Client) Add(a *alert.Alert) {
	a.State = constant.ActiveState
	c.alerts = append(c.alerts, a)
}
