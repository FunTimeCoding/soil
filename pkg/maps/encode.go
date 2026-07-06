package maps

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Encode(m map[string]string) []byte {
	result, e := json.Marshal(m)
	errors.PanicOnError(e)

	return result
}
