package notation

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func MustDecodeBytes(
	value []byte,
	a any,
	verbose bool,
) {
	e := DecodeBytes(value, a)

	if verbose && e != nil {
		fmt.Printf("Could not decode: %s\n", value)
	}

	errors.PanicOnError(e)
}
