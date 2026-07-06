package checksum

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func Parse(input string) map[string]string {
	result := make(map[string]string)

	for _, l := range split.NewLine(input) {
		parts := strings.Split(l, separator.DoubleSpace)
		result[parts[1]] = parts[0]
	}

	return result
}
