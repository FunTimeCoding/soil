package event_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"net/url"
	"strconv"
)

func Path(
	o *query_option.Option,
	page int,
) string {
	q := url.Values{}
	q.Set(constant.PageParameter, strconv.Itoa(page))

	for name, value := range map[string]string{
		constant.Tool:    o.Tool,
		constant.Surface: o.Surface,
		constant.Actor:   o.Actor,
		constant.Kind:    o.Kind,
	} {
		if value != "" {
			q.Set(name, value)
		}
	}

	return (&url.URL{Path: constant.EventsPath, RawQuery: q.Encode()}).String()
}
