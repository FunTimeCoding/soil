package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/chroma"
)

func Collection() {
	c := chroma.NewEnvironment()

	for _, l := range c.Collections() {
		console.Format("Collection: %s\n", l.Name())
	}
}
