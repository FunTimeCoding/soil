package label_change

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func Format(
	key string,
	past string,
	now string,
) string {
	if past == "" {
		return fmt.Sprintf(
			"%s %s%s%s",
			key,
			constant.UnsetMarker,
			constant.ChangeArrow,
			now,
		)
	}

	if now == "" {
		return fmt.Sprintf(
			"%s %s%s %s",
			key,
			past,
			constant.ChangeArrow,
			constant.UnsetMarker,
		)
	}

	return fmt.Sprintf("%s %s%s%s", key, past, constant.ChangeArrow, now)
}
