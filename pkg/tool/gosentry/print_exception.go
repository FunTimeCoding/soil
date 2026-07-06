package gosentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func printException(values []response.ExceptionValue) {
	for _, v := range values {
		fmt.Printf("Exception: %s: %s\n", v.Type, v.Value)
		printStacktrace(v.Stacktrace)
	}
}
