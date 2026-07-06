package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func parseHex(hex string) (int, int, int) {
	var r, g, b int
	_, e := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	errors.PanicOnError(e)

	return r, g, b
}
