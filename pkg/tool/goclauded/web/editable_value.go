package web

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func editableValue(
	kind string,
	metadata map[string]string,
) string {
	if metadata == nil {
		return ""
	}

	switch kind {
	case constant.Complete:
		return metadata[constant.Message]
	case constant.Moment:
		return metadata[constant.Line]
	default:
		return metadata[constant.Body]
	}
}
