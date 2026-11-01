package seed_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
)

func Path(modified bool) string {
	if !modified {
		return constant.DashboardPath
	}

	return fmt.Sprintf(
		"%s?%s=%s",
		constant.DashboardPath,
		constant.SortParameter,
		constant.SortModified,
	)
}
