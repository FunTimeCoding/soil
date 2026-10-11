package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"sync"
)

var IndexDirectory = sync.OnceValue(
	func() string {
		result, e := os.MkdirTemp("", "gosourced-index-")
		errors.PanicOnError(e)

		return result
	},
)
