package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Simple() {
	o := ollama.NewEnvironment()
	console.Line(o.GenerateSimple("What is a car?").Text)
}
