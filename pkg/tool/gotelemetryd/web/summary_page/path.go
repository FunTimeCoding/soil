package summary_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"net/url"
)

func Path(
	window string,
	group string,
) string {
	q := url.Values{}
	q.Set(constant.WindowParameter, window)
	q.Set(constant.GroupParameter, group)

	return (&url.URL{Path: constant.HeatmapPath, RawQuery: q.Encode()}).String()
}
