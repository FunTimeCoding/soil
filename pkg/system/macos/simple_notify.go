package macos

import (
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/soil/pkg/errors"
)

func SimpleNotify(text string) {
	errors.PanicOnError(mack.Notify(text))
}
