package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func editableKey(kind string) string {
	switch kind {
	case constant.Complete:
		return constant.Message
	case constant.Moment:
		return constant.Line
	default:
		return constant.Body
	}
}
