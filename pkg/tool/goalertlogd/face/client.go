package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
)

type Client interface {
	GetAlerts(
		x context.Context,
		p *client.GetAlertsParams,
	) ([]client.AlertsResponse, error)
	GetAlertsAfter(
		x context.Context,
		p *client.GetAlertsAfterParams,
	) ([]client.AlertsResponse, error)
	GetAlertsByIdentifier(
		x context.Context,
		p *client.GetAlertsByIdentifierParams,
	) ([]client.AlertsResponse, error)
	GetRecentAlerts(
		x context.Context,
		p *client.GetRecentAlertsParams,
	) ([]client.AlertsResponse, error)
	GetTopAlerts(
		x context.Context,
		p *client.GetTopAlertsParams,
	) ([]client.TopAlertsResponse, error)
	GetStatus(x context.Context) (*client.StatusResponse, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
