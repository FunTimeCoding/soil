package teams

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
)

func decodeAll[T any](items []json.RawMessage) ([]T, error) {
	result := make([]T, 0, len(items))

	for _, item := range items {
		var v T

		if e := json.Unmarshal(item, &v); e != nil {
			return nil, unexpected.Format("unexpected graph item: %s", e)
		}

		result = append(result, v)
	}

	return result, nil
}
