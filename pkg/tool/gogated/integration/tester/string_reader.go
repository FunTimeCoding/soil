package tester

import (
	"io"
	"strings"
)

func stringReader(s string) io.ReadCloser {
	return io.NopCloser(strings.NewReader(s))
}
