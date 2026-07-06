package macos

import (
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/soil/pkg/errors"
)

func InputDialog(
	subject string,
	body string,
	suggestion string,
) {
	timeout := "5"
	response, e := mack.Dialog(body, subject, suggestion, timeout)
	errors.PanicOnError(e)
	fmt.Printf("%#v", response)
}
