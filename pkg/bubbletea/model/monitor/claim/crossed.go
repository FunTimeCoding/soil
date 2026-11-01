package claim

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func (c *Connection) Crossed(x *edge.Edge) {
	if !x.Down {
		return
	}

	c.updates <- Message{Error: unavailable.Format("%s", x)}
}
