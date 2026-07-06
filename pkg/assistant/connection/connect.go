package connection

import (
	"fmt"
	assistant "github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/gorilla/websocket"
	"net/url"
)

func (c *Connection) Connect() {
	c.Lock()
	defer c.Unlock()
	u := &url.URL{
		Scheme: constant.Socket,
		Host:   fmt.Sprintf("%s:8123", c.host),
		Path:   assistant.Path,
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
		case assistant.AuthenticationRequired:
			errors.PanicOnError(
				c.connection.WriteJSON(
					&authenticateCommand{
						Type:  Authenticate,
						Token: c.token,
					},
				),
			)
		case assistant.AuthenticationInvalid:
			panic("authentication invalid")
		case assistant.AuthenticationSuccess:
			// good
			return
		}
	}
}
