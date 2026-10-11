package delivery

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) Cut(e queue.Entry) string {
	m := d.lookup(e)

	if m == nil || length(e.Body) <= constant.DeliveryBudget {
		return e.Body
	}

	if f := fragment(e.Body, m, constant.DeliveryBudget); f != "" {
		return f
	}

	return d.pointer(m)
}
