package claim

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func (c *Connection) Crossed(x *edge.Edge) {
	if !x.Down {
		return
	}

	c.deliver(message.Message{Error: unavailable.Format("%s", x)})
}
