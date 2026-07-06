package macos

import (
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Beep() {
	errors.PanicOnError(mack.Beep(1))
}
