package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Fast() {
	o := ollama.NewEnvironment()
	r := o.GenerateFast("One short sentence: What is a car?")
	console.Line(r.Text)
	r.Print()
}
