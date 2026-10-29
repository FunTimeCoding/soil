package runtime

import (
	"reflect"
	"runtime"
	"strings"
)

func MethodName(method any) string {
	return strings.TrimSuffix(
		runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(),
		"-fm",
	)
}
