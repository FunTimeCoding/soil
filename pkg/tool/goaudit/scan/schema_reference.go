package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan/open_api"
	"strings"
)

func schemaReference(r open_api.Response) string {
	if r.Content == nil {
		return ""
	}

	j, okay := r.Content["application/json"]

	if !okay {
		return ""
	}

	e := j.Schema.Reference

	if e == "" {
		return ""
	}

	parts := strings.Split(e, "/")

	return parts[len(parts)-1]
}
