package delivery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) messageMinimum(candidates []queue.Entry) int {
	if len(candidates) == 0 {
		return 0
	}

	return min(
		size(
			section(
				constant.DeliveryMessages,
				bodies(candidates, constant.DeliveryIndent),
			),
		),
		size(
			[]string{
				constant.DeliveryMessages,
				fmt.Sprintf(constant.DeliveryIndent, constant.DeliveryExplanation),
			},
		)+d.reserve(candidates),
	)
}
