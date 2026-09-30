package assistant

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/connection"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(
	host string,
	token string,
	l *logger.Logger,
	r face.Reporter,
	o ...Option,
) *Client {
	result := &Client{
		connection: connection.New(host, token),
		recovery:   recovery.New(l, r),
		context:    context.Background(),
	}

	for _, o := range o {
		o(result)
	}

	return result
}
