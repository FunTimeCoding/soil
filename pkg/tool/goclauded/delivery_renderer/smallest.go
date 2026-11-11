package delivery_renderer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) smallest(e queue.Entry) int {
	inline := lineSize(fmt.Sprintf(constant.DeliveryIndent, e.Body))
	m := r.lookup(e)

	if m == nil {
		return inline
	}

	return min(
		inline,
		lineSize(fmt.Sprintf(constant.DeliveryIndent, r.pointer(m))),
	)
}
