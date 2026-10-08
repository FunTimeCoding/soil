package conversations

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func blockLabel(
	role string,
	kind string,
) string {
	if kind == constant.BlockMessage {
		return role
	}

	return kind
}
