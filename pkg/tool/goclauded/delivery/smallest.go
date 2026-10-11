package delivery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) smallest(e queue.Entry) int {
	inline := lineSize(fmt.Sprintf(constant.DeliveryIndent, e.Body))
	m := d.lookup(e)

	if m == nil {
		return inline
	}

	return min(
		inline,
		lineSize(fmt.Sprintf(constant.DeliveryIndent, d.pointer(m))),
	)
}
