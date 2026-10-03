package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gmail"
	"github.com/funtimecoding/soil/pkg/gmail/constant"
)

func main() {
	c := gmail.NewEnvironment().Load()
	r := c.Unread()
	console.Format("Unread (%d):\n", len(r.Messages))

	for _, m := range r.Messages {
		var subject string

		for _, h := range c.Message(m).Payload.Headers {
			if h.Name == constant.SubjectHeader {
				subject = h.Value

				break
			}
		}

		console.Format("- %s\n", subject)
	}
}
