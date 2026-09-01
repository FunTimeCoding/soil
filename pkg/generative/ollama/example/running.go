package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Running() {
	o := ollama.NewEnvironment()

	for _, m := range o.MustRunning() {
		console.Format("Running: %+v\n", m)
	}
}
