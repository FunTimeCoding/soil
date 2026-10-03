package macos

import (
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Alert(
	subject string,
	body string,
) {
	icon := "informational"
	timeout := "5"
	response, e := mack.Alert(subject, body, icon, timeout)
	errors.PanicOnError(e)
	fmt.Printf("%#v", response)
}
