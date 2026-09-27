package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strconv"
)

func streamOverride(r *http.Request) *uint {
	raw := r.Header.Get(webConstant.LastEvent)

	if raw == "" {
		raw = r.URL.Query().Get(constant.After)
	}

	if raw == "" {
		return nil
	}

	parsed, e := strconv.ParseUint(raw, 10, 64)

	if e != nil {
		return nil
	}

	result := uint(parsed)

	return &result
}
