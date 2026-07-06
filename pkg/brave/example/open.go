package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/generative/openai/site"
)

func Open() {
	a := argument.NewSimple("brave-open")
	a.ParseSimple()
	site.New().OpenChat(a.RequiredPositional(0, "CHAT"))
}
