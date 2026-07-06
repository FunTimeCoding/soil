package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/open_webui"
)

func Load() {
	c := open_webui.NewEnvironment()
	fmt.Printf("Folders: %s\n", c.Folders())

	if false {
		fmt.Printf("Functions: %s\n", c.Functions())
		fmt.Printf("Users: %s\n", c.Users())
		fmt.Printf("Files: %s\n", c.Files())
		fmt.Printf("Memories: %s\n", c.Memories())
		fmt.Printf("Knowledge: %s\n", c.Knowledge())
		fmt.Printf("Models: %s\n", c.Models())
		fmt.Printf("Chats: %s\n", c.Chats())
		fmt.Printf("Memories: %s\n", c.Memories())
		fmt.Printf("Notes: %s\n", c.Notes())
		fmt.Printf("Prompts: %s\n", c.Prompts())
	}
}
