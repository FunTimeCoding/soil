package embed

import "github.com/funtimecoding/soil/pkg/face"

func Single(
	e face.Embedder,
	v string,
) ([]float32, error) {
	result, f := e.Embed([]string{v})

	if f != nil {
		return nil, f
	}

	return result[0], nil
}
