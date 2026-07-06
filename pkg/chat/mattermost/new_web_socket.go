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
) *model.WebSocketClient {
	result, e := model.NewWebSocketClient4(
		locator.New(host).Scheme(constant.SecureSocket).String(),
		token,
	)
	errors.PanicOnError(e)

	return result
}
