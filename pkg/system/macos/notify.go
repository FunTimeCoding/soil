package macos

import (
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Notify(
	title string,
	subtitle string,
	body string,
) {
	errors.PanicOnError(mack.Notify(body, title, subtitle))
}
