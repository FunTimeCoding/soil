package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func eventsLink(
	o *store.QueryOption,
	page int,
) string {
	link := fmt.Sprintf("/events?page=%d", page)

	if o.Tool != "" {
		link = fmt.Sprintf("%s&tool=%s", link, o.Tool)
	}

	if o.Surface != "" {
		link = fmt.Sprintf("%s&surface=%s", link, o.Surface)
	}

	if o.Actor != "" {
		link = fmt.Sprintf("%s&actor=%s", link, o.Actor)
	}

	if o.Kind != "" {
		link = fmt.Sprintf("%s&kind=%s", link, o.Kind)
	}

	return link
}
