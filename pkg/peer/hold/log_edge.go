package hold

import "github.com/funtimecoding/soil/pkg/peer/reacher/edge"

func (h *Hold) logEdge(x *edge.Edge) {
	h.logger.Plain("%s", x)
}
