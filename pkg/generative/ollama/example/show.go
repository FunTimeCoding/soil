package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Show() {
	o := ollama.NewEnvironment()
	console.Format("Show: %+v\n", o.MustShow(constant.Llama31))
}
