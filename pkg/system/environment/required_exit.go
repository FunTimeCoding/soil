package environment

import (
	"github.com/funtimecoding/soil/pkg/system"
	"os"
)

func RequiredExit(
	name string,
	code int,
) string {
	result := os.Getenv(name)

	if result == "" {
		system.Exitf(code, "%s not set\n", name)
	}

	return result
}
