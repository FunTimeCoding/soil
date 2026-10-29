package unit

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func sentModel(r *http.Request) string {
	var result sentBody
	notation.MustDecodeBytes(system.ReadAll(r.Body), &result, false)

	return result.Model
}
