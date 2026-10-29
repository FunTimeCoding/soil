package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
)

type Client interface {
	GetLogs(
		x context.Context,
		p *client.GetLogsParams,
	) (client.LogsResponse, error)
	GetReports(x context.Context) ([]client.ReportResponse, error)
	Reachable(x context.Context) error
	Close()
}
