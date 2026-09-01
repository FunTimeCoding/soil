package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/open_webui"
)

func Load() {
	c := open_webui.NewEnvironment()
	console.Format("Folders: %s\n", c.Folders())

	if false {
		console.Format("Functions: %s\n", c.Functions())
		console.Format("Users: %s\n", c.Users())
		console.Format("Files: %s\n", c.Files())
		console.Format("Memories: %s\n", c.Memories())
		console.Format("Knowledge: %s\n", c.Knowledge())
		console.Format("Models: %s\n", c.Models())
		console.Format("Chats: %s\n", c.Chats())
		console.Format("Memories: %s\n", c.Memories())
		console.Format("Notes: %s\n", c.Notes())
		console.Format("Prompts: %s\n", c.Prompts())
	}
}
