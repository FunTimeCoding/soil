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
			&Key{Type: constant.LadybirdKeyDown, Value: string(r)},
			&Key{Type: constant.LadybirdKeyUp, Value: string(r)},
		)
	}

	return &Actions{
		Context: context,
		Actions: []*KeyActions{
			{
				Type:       constant.LadybirdKeySource,
				Identifier: constant.LadybirdKeyboardSource,
				Actions:    keys,
			},
		},
	}
}
