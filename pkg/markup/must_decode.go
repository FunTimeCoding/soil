package markup

import "github.com/funtimecoding/soil/pkg/errors"

func MustDecode(
	value string,
	structure any,
) {
	errors.PanicOnError(Decode(value, structure))
}
