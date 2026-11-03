package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
)

func answerOf(
	v any,
	e error,
) string {
	if e != nil {
		if c := classify.Class(e); c != "" {
			return c
		}

		return e.Error()
	}

	b, f := json.Marshal(v)
	errors.PanicOnError(f)

	return string(b)
}
