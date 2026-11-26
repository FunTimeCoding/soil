package ladybird

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
)

func decodeResult[T any](
	raw json.RawMessage,
	e error,
) (*T, error) {
	if e != nil {
		return nil, e
	}

	result := new(T)

	if f := notation.DecodeBytes(raw, result); f != nil {
		return nil, fmt.Errorf("decode result: %w", f)
	}

	return result, nil
}
