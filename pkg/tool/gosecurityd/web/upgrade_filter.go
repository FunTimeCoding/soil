package web

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"net/http"
	"strconv"
)

func upgradeFilter(r *http.Request) (string, bool, error) {
	q := r.URL.Query()
	v := q.Get(constant.KernelParameter)

	if v == "" {
		return q.Get(constant.ReleaseParameter), false, nil
	}

	kernel, e := strconv.ParseBool(v)

	if e != nil {
		return "", false, validation.New("kernel not true or false: %s", v)
	}

	return q.Get(constant.ReleaseParameter), kernel, nil
}
