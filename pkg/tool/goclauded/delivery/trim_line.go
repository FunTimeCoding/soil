package delivery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func trimLine(count int) string {
	return fmt.Sprintf(
		constant.DeliveryIndent,
		fmt.Sprintf(constant.DeliveryMemoryTrim, count),
	)
}
