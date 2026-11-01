package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/trim"
	"strings"
)

func failed(
	output string,
	stderr string,
	failure error,
) string {
	var lines []string

	if output != "" {
		lines = append(lines, output, "")
	}

	lines = append(lines, fmt.Sprintf("error: %v", failure))

	if s := trim.NewLine(stderr); s != "" {
		lines = append(lines, s)
	}

	return strings.Join(lines, "\n")
}
