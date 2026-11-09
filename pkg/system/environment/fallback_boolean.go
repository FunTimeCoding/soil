package environment

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"strconv"
)

func FallbackBoolean(
	name string,
	fallback bool,
) bool {
	result := os.Getenv(name)

	if result == "" {
		return fallback
	}

	v, e := strconv.ParseBool(result)
	errors.PanicOnError(e)

	return v
}
