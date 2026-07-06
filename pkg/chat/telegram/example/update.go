package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
)

func Update() {
	c := telegram.NewEnvironment()
	f := constant.Format

	for _, m := range c.Messages() {
		fmt.Println(m.Format(f))
	}
}
