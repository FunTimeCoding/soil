package example

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
	"github.com/funtimecoding/soil/pkg/console"
)

func Update() {
	c := telegram.NewEnvironment()
	f := constant.TelegramFormat

	for _, m := range c.Messages() {
		console.Line(m.Format(f))
	}
}
