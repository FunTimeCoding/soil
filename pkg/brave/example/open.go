package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/generative/openai/site"
)

func Open() {
	a := argument.NewSimple("brave-open")
	a.ParseSimple()
	s := site.New()
	defer s.Close()
	s.OpenChat(a.RequiredPositional(0, "CHAT"))
}
