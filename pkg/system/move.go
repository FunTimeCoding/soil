package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"syscall"
)

func Move(
	from string,
	to string,
) {
	if e := os.Rename(from, to); e != nil {
		if errors.Is(e, syscall.EXDEV) {
			MoveCopy(from, to)
		} else {
			errors.PanicOnError(e)
		}
	}
}
