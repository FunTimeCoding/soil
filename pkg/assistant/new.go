package assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/connection"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/hold"
)

func New(
	host string,
	token string,
	l *logger.Logger,
	r face.Reporter,
	o ...Option,
) *Client {
	result := &Client{recovery: recovery.New(l, r)}

	for _, o := range o {
		o(result)
	}

	result.connection = connection.New(host, token, result.dispatch)
	result.hold = hold.New(host, result.connection, l, r)

	return result
}
