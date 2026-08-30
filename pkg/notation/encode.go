package notation

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Encode(
	a any,
	indent bool,
) string {
	options := []json.Options{json.Deterministic(true)}

	if indent {
		options = append(options, jsontext.WithIndent("    "))
	}

	b, e := json.Marshal(a, options...)
	errors.PanicOnError(e)

	return string(b)
}
