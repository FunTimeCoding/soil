package web

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"net/http"
	"strconv"
)

func minimum(
	r *http.Request,
	fallback float64,
) (float64, error) {
	v := r.URL.Query().Get(constant.ScoreField)

	if v == "" {
		return fallback, nil
	}

	result, e := strconv.ParseFloat(v, 64)

	if e != nil {
		return 0, validation.New("score not a number: %s", v)
	}

	return result, nil
}
