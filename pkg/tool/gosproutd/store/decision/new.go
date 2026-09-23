package decision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/choice"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/frame"
)

func New(
	session string,
	question string,
	defaultAction string,
	choices []*choice.Choice,
	frames []*frame.Frame,
) *Decision {
	return &Decision{
		Session:       session,
		Question:      question,
		DefaultAction: defaultAction,
		State:         constant.StateOpen,
		Choices:       choices,
		Frames:        frames,
	}
}
