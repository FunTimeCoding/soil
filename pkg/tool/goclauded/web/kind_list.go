package web

import (
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func kindList(r *http.Request) []string {
	raw := r.URL.Query().Get(constant.Kinds)

	if raw == "" {
		return nil
	}

	return strings.DeleteEmpty(split.Comma(raw))
}
