package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
)

func Echo() {
	c := telegram.NewEnvironment()
	f := constant.TelegramFormat

	for _, m := range c.Messages() {
		fmt.Println(m.Format(f))
		c.Reply(m.Update, m.Text)
	}
}
