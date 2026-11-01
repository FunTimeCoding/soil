package assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/connection"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/peer/hold"
)

type Client struct {
	connection *connection.Connection
	hold       *hold.Hold
	recovery   *recovery.Recovery
	subscriber connection.Subscriber
}
