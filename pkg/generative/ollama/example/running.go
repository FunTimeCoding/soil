package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Running() {
	o := ollama.NewEnvironment()

	for _, m := range o.MustRunning() {
		fmt.Printf("Running: %+v\n", m)
	}
}
