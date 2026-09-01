package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func List() {
	o := ollama.NewEnvironment()

	for _, m := range o.MustList() {
		console.Format("Model: %+v\n", m)
	}
}
