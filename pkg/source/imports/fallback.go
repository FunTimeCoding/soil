package imports

import "github.com/funtimecoding/soil/pkg/strings/join"

func Fallback(
	subsystem string,
	declaredName string,
) string {
	return join.Empty(subsystem, upperFirst(declaredName))
}
