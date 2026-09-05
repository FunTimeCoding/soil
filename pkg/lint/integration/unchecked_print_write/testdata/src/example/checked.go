package example

import (
	"fmt"
	"os"
)

func Checked() {
	_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", "something")
}
