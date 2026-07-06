package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gmail"
	"github.com/funtimecoding/soil/pkg/gmail/constant"
)

func main() {
	// TODO: how to run multiple in sequence and verify all are logged in?
	//  With --account $name
	//  If no token exists or token is invalid: Output that as single element
	//   Maybe a Brave profile with the name can be opened?
	c := gmail.NewEnvironment().Load()
	r := c.Unread()
	fmt.Printf("Unread (%d):\n", len(r.Messages))

	for _, m := range r.Messages {
		var subject string

		for _, h := range c.Message(m).Payload.Headers {
			if h.Name == constant.SubjectHeader {
				subject = h.Value

				break
			}
		}

		fmt.Printf("- %s\n", subject)
	}
}
