package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
)

func Update() {
	c := telegram.NewEnvironment()
	f := constant.TelegramFormat

	for _, m := range c.Messages() {
		fmt.Println(m.Format(f))
	}
}
