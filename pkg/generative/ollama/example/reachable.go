package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
)

func Reachable() {
	o := ollama.NewEnvironment()
	defer o.Close()
	errors.PanicOnError(o.Reachable(context.Background()))
}
