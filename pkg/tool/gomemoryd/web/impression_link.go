package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strconv"
)

func impressionLink(days int) string {
	if days == 7 {
		return constant.ImpressionsPath
	}

	return fmt.Sprintf(
		"%s?days=%s",
		constant.ImpressionsPath,
		strconv.Itoa(days),
	)
}
