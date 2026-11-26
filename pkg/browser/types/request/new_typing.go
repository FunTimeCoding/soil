package request

import "github.com/funtimecoding/soil/pkg/browser/constant"

func NewTyping(
	context string,
	text string,
) *Actions {
	var keys []*Key

	for _, r := range text {
		keys = append(
			keys,
			&Key{Type: constant.KeyDown, Value: string(r)},
			&Key{Type: constant.KeyUp, Value: string(r)},
		)
	}

	return &Actions{
		Context: context,
		Actions: []*KeyActions{
			{
				Type:       constant.KeySource,
				Identifier: constant.KeyboardSource,
				Actions:    keys,
			},
		},
	}
}
