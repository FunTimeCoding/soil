package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/chat/discord"
	"github.com/funtimecoding/soil/pkg/system"
)

func DeleteLoop() {
	a := argument.NewSimple("discord-delete-loop")
	a.String(
		DeleteLoopArgument,
		"",
		"Delete messages in a loop, requires channel",
	)
	a.ParseSimple()
	c := discord.NewEnvironment()
	c.Listen(false)
	defer c.Close()

	if h := a.GetString(DeleteLoopArgument); h != "" {
		c.DeleteLoop(h)

		return
	}

	if false {
		c.PrintStatus()
	}

	system.KillSignalBlock()
}
