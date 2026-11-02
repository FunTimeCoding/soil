package unit

import (
	"bufio"
	"strings"
	"testing"
)

func awaitLine(
	t *testing.T,
	r *bufio.Reader,
	substring string,
) {
	t.Helper()

	for {
		line, e := r.ReadString('\n')

		if e != nil {
			t.Fatalf("stream ended before a line with %q: %v", substring, e)
		}

		if strings.Contains(line, substring) {
			return
		}
	}
}
