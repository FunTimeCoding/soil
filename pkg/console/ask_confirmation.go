package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/reader"
	"strings"
)

func AskConfirmation(s string) bool {
	r := reader.New()

	for {
		fmt.Printf("%s [y/N]: ", s)
		l := r.Line()
		l = strings.ToLower(strings.TrimSpace(l))

		switch l {
		case "y", "yes":
			return true
		default:
			return false
		}
	}
}
