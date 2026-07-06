package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func List() {
	o := ollama.NewEnvironment()

	for _, m := range o.MustList() {
		fmt.Printf("Model: %+v\n", m)
	}
}
