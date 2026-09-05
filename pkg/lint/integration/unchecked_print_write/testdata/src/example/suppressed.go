package example

import (
	"fmt"
	"os"
)

func Suppressed() {
	fmt.Fprintf(os.Stderr, "error: %s\n", "something") // goanalyze:ignore
}
