package example

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram"
	"github.com/funtimecoding/soil/pkg/console"
)

func User() {
	c := telegram.NewEnvironment()
	console.Format("User: %s\n", c.Self().UserName)
}
