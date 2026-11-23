package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"strconv"
)

func pageSize(limit int) (int, string) {
	if limit <= 0 {
		limit = constant.TeamsDefaultLimit
	}

	return limit, strconv.Itoa(min(limit, microsoftConstant.PageMaximum))
}
