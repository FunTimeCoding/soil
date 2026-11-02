package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
)

type Client interface {
	GetEvents(
		x context.Context,
		p *client.GetEventsParams,
	) ([]client.EventEntry, error)
	GetSummary(
		x context.Context,
		p *client.GetSummaryParams,
	) ([]client.SummaryEntry, error)
	PostEvent(
		x context.Context,
		b client.PostEventJSONRequestBody,
	) (*client.EventResponse, error)
	Reachable(x context.Context) error
	Close()
}
