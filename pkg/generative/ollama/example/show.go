package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
)

func Show() {
	o := ollama.NewEnvironment()
	fmt.Printf("Show: %+v\n", o.MustShow(constant.Llama31))
}
