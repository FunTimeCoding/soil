package notation

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"strings"
)

func MustDecodeStream[T any](value string) []*T {
	d := json.NewDecoder(strings.NewReader(value))
	var result []*T

	for d.More() {
		v := new(T)
		errors.PanicOnError(d.Decode(v))
		result = append(result, v)
	}

	return result
}
