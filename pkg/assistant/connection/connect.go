package connection

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/gorilla/websocket"
	"net/url"
)

func (c *Connection) Connect() {
	c.Lock()
	defer c.Unlock()
	u := &url.URL{
		Scheme: webConstant.Socket,
		Host:   fmt.Sprintf("%s:8123", c.host),
		Path:   constant.Path,
	}
	var e error
	c.connection, _, e = websocket.DefaultDialer.Dial(
		u.String(),
		nil,
	)
	errors.PanicOnError(e)

	for {
		var m message.Message
		errors.PanicOnError(c.connection.ReadJSON(&m))

		switch m.Type {
		case constant.AuthenticationRequired:
			errors.PanicOnError(
				c.connection.WriteJSON(
					&authenticateCommand{
						Type:  constant.Authenticate,
						Token: c.token,
					},
				),
			)
		case constant.AuthenticationInvalid:
			panic("authentication invalid")
		case constant.AuthenticationSuccess:
			return
		}
	}
}
