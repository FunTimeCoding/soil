package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/embed"
)

func Target() {
	e := embed.NewEnvironment()
	result, f := e.Embed([]string{"What are embeddings?", "A second document."})
	errors.PanicOnError(f)

	for _, vector := range result {
		console.Format(
			"Dimensions: %d, first values: %v\n",
			len(vector),
			vector[:4],
		)
	}
}
