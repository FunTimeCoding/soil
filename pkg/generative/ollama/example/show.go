package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Show() {
	o := ollama.NewEnvironment()
	fmt.Printf("Show: %+v\n", o.MustShow(constant.Llama31))
}
