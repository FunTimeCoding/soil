package process

import (
	"errors"
	"golang.org/x/sys/unix"
	"os"
)

func Exited(e error) bool {
	return errors.Is(e, unix.ESRCH) || errors.Is(e, os.ErrProcessDone)
}
