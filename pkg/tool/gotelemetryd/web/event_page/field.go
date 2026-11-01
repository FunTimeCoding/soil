package event_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
)

func field(
	o *query_option.Option,
	name string,
) string {
	switch name {
	case constant.Tool:
		return o.Tool
	case constant.Surface:
		return o.Surface
	case constant.Actor:
		return o.Actor
	case constant.Kind:
		return o.Kind
	}

	return ""
}
