package delivery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (d *Delivery) messageLines(
	candidates []queue.Entry,
	room int,
) []string {
	whole := section(
		constant.DeliveryMessages,
		bodies(candidates, constant.DeliveryIndent),
	)

	if size(whole) <= room {
		return whole
	}

	result := []string{
		constant.DeliveryMessages,
		fmt.Sprintf(constant.DeliveryIndent, constant.DeliveryExplanation),
	}
	remaining := room - size(result)

	for i, e := range candidates {
		line := d.place(e, remaining-d.reserve(candidates[i+1:]))
		result = append(result, line)
		remaining -= lineSize(line)
	}

	return result
}
