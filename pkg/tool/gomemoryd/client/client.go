package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
	"time"
)

type Client interface {
	face.Reachable
	Address() string
	VersionsSince(
		x context.Context,
		since string,
		limit int,
	) ([]version_entry.Entry, error)
	SaveImpression(
		x context.Context,
		content string,
		source string,
		capturedAt time.Time,
	) error
	RedactedMemories(x context.Context) (map[int64]bool, error)
	Statistics(x context.Context) (*client.Statistics, error)
	Relations(x context.Context) ([]client.Relation, error)
}
