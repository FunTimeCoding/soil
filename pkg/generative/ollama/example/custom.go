package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
)

func Custom() {
	o := ollama.NewEnvironment(ollama.WithSecure(true))
	console.Format("Version: %s\n", o.MustVersion())
	r := o.MustGenerate(
		generate_request.New().Prompt("One short sentence: What is a car?").Model(
			"codellama:7b",
		),
	)
	console.Line(r.Text)
	r.Print()
}
