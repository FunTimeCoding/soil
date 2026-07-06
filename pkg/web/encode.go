package web

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func Encode(
	w http.ResponseWriter,
	a any,
) {
	errors.PanicOnError(json.NewEncoder(w).Encode(a))
}
