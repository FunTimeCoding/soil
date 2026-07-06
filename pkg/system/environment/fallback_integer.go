package environment

import (
	"github.com/funtimecoding/soil/pkg/strings"
	"os"
)

func FallbackInteger(
	name string,
	fallback int,
) int {
	result := os.Getenv(name)

	if result == "" {
		return fallback
	}

	return strings.MustToInteger(result)
}
