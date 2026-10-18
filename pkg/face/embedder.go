package face

import "context"

type Embedder interface {
	Embed(
		x context.Context,
		v []string,
	) ([][]float32, error)
}
