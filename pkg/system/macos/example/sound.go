package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/sound"
	"github.com/funtimecoding/soil/pkg/system"
)

func Sound() {
	fmt.Println(system.Run(sound.Afplay, sound.SosumiPath))
	fmt.Println(system.Run(sound.Afplay, sound.TinkPath))
}
