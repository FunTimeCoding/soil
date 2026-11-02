package mattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/gorilla/websocket"
	"github.com/mattermost/mattermost/server/public/model"
	"net"
)

func newWebSocket(
	x context.Context,
	host string,
	token string,
	insecure bool,
) (*model.WebSocketClient, error) {
	scheme := constant.SecureSocket

	if insecure {
		scheme = constant.Socket
	}

	d := *websocket.DefaultDialer
	d.NetDialContext = func(
		_ context.Context,
		network string,
		address string,
	) (net.Conn, error) {
		var n net.Dialer

		return n.DialContext(x, network, address)
	}

	return model.NewWebSocketClient4WithDialer(
		&d,
		locator.New(host).Scheme(scheme).String(),
		token,
	)
}
