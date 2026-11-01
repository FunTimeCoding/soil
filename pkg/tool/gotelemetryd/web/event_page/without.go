package event_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
)

func without(
	o *query_option.Option,
	name string,
) *query_option.Option {
	result := *o

	switch name {
	case constant.Tool:
		result.Tool = ""
	case constant.Surface:
		result.Surface = ""
	case constant.Actor:
		result.Actor = ""
	case constant.Kind:
		result.Kind = ""
	}

	return &result
}
