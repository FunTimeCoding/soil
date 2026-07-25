package mattermost

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func newWebSocket(
	host string,
	token string,
	insecure bool,
) *model.WebSocketClient {
	scheme := constant.SecureSocket

	if insecure {
		scheme = constant.Socket
	}

	result, e := model.NewWebSocketClient4(
		locator.New(host).Scheme(scheme).String(),
		token,
	)
	errors.PanicOnError(e)

	return result
}
