package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
)

func Update() {
	c := telegram.NewEnvironment()
	f := telegram.Format

	for _, m := range c.Messages() {
		fmt.Println(m.Format(f))
	}
}
