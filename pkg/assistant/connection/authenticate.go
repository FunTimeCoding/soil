package connection

import (
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/assistant/types/authenticate_command"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
)

func (c *Connection) authenticate() error {
	for {
		var m message.Message

		if e := c.connection.ReadJSON(&m); e != nil {
			return e
		}

		switch m.Type {
		case constant.AuthenticationRequired:
			if e := c.connection.WriteJSON(
				authenticate_command.New(constant.Authenticate, c.token),
			); e != nil {
				return e
			}
		case constant.AuthenticationInvalid:
			return forbidden.Format("%s", constant.AuthenticationInvalid)
		case constant.AuthenticationSuccess:
			return nil
		}
	}
}
