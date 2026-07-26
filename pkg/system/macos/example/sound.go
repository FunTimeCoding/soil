package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/sound/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Sound() {
	fmt.Println(system.Run(constant.Afplay, constant.SosumiPath))
	fmt.Println(system.Run(constant.Afplay, constant.TinkPath))
}
