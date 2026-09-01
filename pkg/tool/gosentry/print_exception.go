package gosentry

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func printException(values []response.ExceptionValue) {
	for _, v := range values {
		console.Format("Exception: %s: %s\n", v.Type, v.Value)
		printStacktrace(v.Stacktrace)
	}
}
