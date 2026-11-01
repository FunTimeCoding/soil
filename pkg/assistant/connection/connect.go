package connection

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/gorilla/websocket"
	"net"
	"net/url"
)

func (c *Connection) connect(x context.Context) error {
	c.Lock()
	defer c.Unlock()
	host := c.host

	if _, _, f := net.SplitHostPort(host); f != nil {
		host = net.JoinHostPort(host, constant.Port)
	}

	u := &url.URL{Scheme: webConstant.Socket, Host: host, Path: constant.Path}
	s, r, e := websocket.DefaultDialer.DialContext(x, u.String(), nil)

	if r != nil {
		defer errors.LogClose(r.Body)
	}

	if e != nil {
		if r != nil {
			return refusal.New(constant.ConnectOperation, r, nil)
		}

		return e
	}

	c.connection = s
	c.subscribers = make(map[uint64]Subscriber)
	c.lastIdentifier = 0

	return c.authenticate()
}
