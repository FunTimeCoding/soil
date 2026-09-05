package example

import (
	"fmt"
	"os"
)

func UncheckedStdout() {
	fmt.Fprintf(
		os.Stdout,
		"output: %s\n",
		"something",
	)
}
