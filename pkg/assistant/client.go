package assistant

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/connection"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
)

type Client struct {
	connection *connection.Connection
	recovery   *recovery.Recovery
	context    context.Context
	subscriber connection.Subscriber
}
