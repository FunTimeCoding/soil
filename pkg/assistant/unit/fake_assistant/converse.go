package fake_assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/connection/subscribe_command"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/assistant/types/authenticate_command"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gorilla/websocket"
)

func (s *Server) converse(
	c *websocket.Conn,
	invalid bool,
) {
	defer errors.LogClose(c)

	if say(c, `{"type":"auth_required"}`) != nil {
		return
	}

	var a authenticate_command.Command

	if c.ReadJSON(&a) != nil {
		return
	}

	if invalid {
		errors.LogOnError(say(c, `{"type":"auth_invalid"}`))

		return
	}

	if say(c, `{"type":"auth_ok"}`) != nil {
		return
	}

	var b subscribe_command.SubscribeCommand

	if c.ReadJSON(&b) != nil {
		return
	}

	s.mutex.Lock()
	s.socket = c
	s.identifier = b.Identifier
	s.subscription++
	s.mutex.Unlock()

	for {
		var m message.Message

		if c.ReadJSON(&m) != nil {
			return
		}
	}
}
