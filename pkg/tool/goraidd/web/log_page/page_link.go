package log_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
)

func pageLink(
	offset int,
	start string,
	end string,
) string {
	link := fmt.Sprintf(
		"%s?%s=%d",
		constant.LogsPath,
		constant.OffsetParameter,
		offset,
	)

	if start != "" {
		link = fmt.Sprintf("%s&%s=%s", link, constant.StartParameter, start)
	}

	if end != "" {
		link = fmt.Sprintf("%s&%s=%s", link, constant.EndParameter, end)
	}

	return link
}
