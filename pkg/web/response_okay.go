package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"slices"
)

func ResponseOkay(r *http.Response) bool {
	return slices.Contains(constant.OkayStatusCodes, r.StatusCode)
}
