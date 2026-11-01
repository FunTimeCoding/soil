package impression_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func Path(days int) string {
	if days == constant.DefaultImpressionDays {
		return constant.ImpressionsPath
	}

	return fmt.Sprintf(
		"%s?%s=%d",
		constant.ImpressionsPath,
		constant.Days,
		days,
	)
}
