package notation

import (
	"encoding/json/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Marshal(a any) []byte {
	result, e := json.Marshal(a, json.Deterministic(true))
	errors.PanicOnError(e)

	return result
}
