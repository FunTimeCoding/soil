package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"sync"
)

var indexDirectory = sync.OnceValue(
	func() string {
		result, e := os.MkdirTemp("", "gosourced-index-")
		errors.PanicOnError(e)

		return result
	},
)
