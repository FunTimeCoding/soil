package delivery_renderer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) place(
	e queue.Entry,
	available int,
) string {
	inline := fmt.Sprintf(constant.DeliveryIndent, e.Body)

	if lineSize(inline) <= available {
		return inline
	}

	m := r.lookup(e)

	if m == nil {
		return inline
	}

	if cut := fragment(
		e.Body,
		m,
		available-lineSize(fmt.Sprintf(constant.DeliveryIndent, "")),
	); cut != "" {
		return fmt.Sprintf(constant.DeliveryIndent, cut)
	}

	return fmt.Sprintf(constant.DeliveryIndent, r.pointer(m))
}
