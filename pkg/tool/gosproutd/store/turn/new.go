package turn

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"

func New(
	decisionIdentifier uint,
	author constant.Author,
	content string,
) *Turn {
	return &Turn{
		DecisionIdentifier: decisionIdentifier,
		Author:             author,
		Content:            content,
	}
}
