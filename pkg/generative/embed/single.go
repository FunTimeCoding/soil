package embed

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
)

func Single(
	x context.Context,
	e face.Embedder,
	v string,
) ([]float32, error) {
	result, f := e.Embed(x, []string{v})

	if f != nil {
		return nil, f
	}

	return result[0], nil
}
