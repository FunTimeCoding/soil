package delivery_renderer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (r *Renderer) messageMinimum(candidates []queue.Entry) int {
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
		)+r.reserve(candidates),
	)
}
