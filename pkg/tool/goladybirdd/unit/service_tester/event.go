package service_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (o *Tester) Event(
	method string,
	context string,
	locator string,
) *message.Reply {
	return message.NewEvent(
		method,
		notation.Encode(
			map[string]any{"context": context, "url": locator},
			false,
		),
	)
}
