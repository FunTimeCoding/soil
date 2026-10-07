package mattermost

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func newWebSocket(
	host string,
	token string,
	insecure bool,
) (*model.WebSocketClient, error) {
	scheme := constant.SecureSocket

	if insecure {
		scheme = constant.Socket
	}

	return model.NewWebSocketClient4(
		locator.New(host).Scheme(scheme).String(),
		token,
	)
}
