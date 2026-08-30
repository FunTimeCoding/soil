package notation

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func MarshalIndentBytes(a any) []byte {
	result, e := json.Marshal(
		a,
		json.Deterministic(true),
		jsontext.WithIndent("\t"),
	)
	errors.PanicOnError(e)

	return result
}
