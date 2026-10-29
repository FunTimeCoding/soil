package web

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"net/http"
)

func identifierFromQuery(r *http.Request) (uint, error) {
	v := r.URL.Query().Get(constant.Identifier)

	if result := strings.ToUnsignedInteger(v, 0); result != 0 {
		return result, nil
	}

	return 0, validation.New("invalid entry: %q", v)
}
