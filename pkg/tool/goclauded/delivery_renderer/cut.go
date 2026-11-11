package delivery_renderer

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) Cut(e queue.Entry) string {
	m := r.lookup(e)

	if m == nil || length(e.Body) <= constant.DeliveryBudget {
		return e.Body
	}

	if f := fragment(e.Body, m, constant.DeliveryBudget); f != "" {
		return f
	}

	return r.pointer(m)
}
