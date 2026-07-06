package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/chroma"
)

func Collection() {
	c := chroma.NewEnvironment()

	for _, l := range c.Collections() {
		fmt.Printf("Collection: %s\n", l.Name())
	}
}
