package delivery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) place(
	e queue.Entry,
	available int,
) string {
	inline := fmt.Sprintf(constant.DeliveryIndent, e.Body)

	if lineSize(inline) <= available {
		return inline
	}

	m := d.lookup(e)

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

	return fmt.Sprintf(constant.DeliveryIndent, d.pointer(m))
}
