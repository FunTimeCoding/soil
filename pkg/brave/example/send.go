package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/generative/openai/site"
)

func Send() {
	a := argument.NewSimple("brave-send")
	a.ParseSimple()
	site.New().Send(a.RequiredPositional(0, "TEXT"))
}
