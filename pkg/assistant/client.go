package assistant

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/connection"
	"github.com/funtimecoding/soil/pkg/face"
)

type Client struct {
	connection *connection.Connection
	reporter   face.Reporter
	context    context.Context
	subscriber connection.Subscriber
}
