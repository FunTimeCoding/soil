package example

import (
	"fmt"
	"os"
)

func Unchecked() {
	fmt.Fprintf(
		os.Stderr,
		"error: %s\n",
		"something",
	)
}
