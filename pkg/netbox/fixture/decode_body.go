package fixture

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func decodeBody(q *http.Request) map[string]any {
	var result map[string]any
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&result))

	return result
}
