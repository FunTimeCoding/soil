package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"time"
)

func (c *Client) SeedAlert(
	name string,
	labels map[string]string,
	start time.Time,
) client.AlertsResponse {
	c.identifier++
	result := client.AlertsResponse{
		Identifier:  c.identifier,
		Fingerprint: name,
		Name:        name,
		Labels:      labels,
		Start:       start.Format(constant.DateFormat),
		Status:      client.Firing,
	}
	c.alerts = append(c.alerts, result)

	return result
}
