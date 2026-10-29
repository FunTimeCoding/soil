package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"time"
)

type Client interface {
	PostGenerate(
		x context.Context,
		files []string,
		date *time.Time,
	) (client.GenerateResponse, error)
	GetStatus(x context.Context) (client.StatusResponse, error)
	Reachable(x context.Context) error
	Close()
}
