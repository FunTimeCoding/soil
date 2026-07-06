package notation

import (
	"bytes"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Encode(
	a any,
	indent bool,
) string {
	b := &bytes.Buffer{}
	e := json.NewEncoder(b)
	e.SetEscapeHTML(false)

	if indent {
		e.SetIndent("", "    ")
	}

	errors.PanicOnError(e.Encode(a))

	return strings.TrimSuffix(b.String(), separator.Unix)
}
