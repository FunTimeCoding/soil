package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
)

func User() {
	c := telegram.NewEnvironment()
	fmt.Printf("User: %s\n", c.Self().UserName)
}
