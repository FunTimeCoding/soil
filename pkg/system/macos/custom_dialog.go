package macos

import (
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/soil/pkg/errors"
)

func CustomDialog(
	subject string,
	body string,
) {
	timeout := 5
	response, e := mack.DialogBox(
		mack.DialogOptions{
			Title:    subject,
			Text:     body,
			Buttons:  "A, B, C",
			Duration: timeout,
		},
	)
	errors.PanicOnError(e)
	fmt.Printf("%#v", response)
}
