package web

import (
	"encoding/json/v2"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func Encode(
	w http.ResponseWriter,
	a any,
) {
	errors.PanicOnError(json.MarshalWrite(w, a, json.Deterministic(true)))
}
