package goquery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func lineRanges(lines []int) string {
	var parts []string

	for i := 0; i < len(lines); {
		j := i

		for j+1 < len(lines) && lines[j+1] == lines[j]+1 {
			j++
		}

		if i == j {
			parts = append(parts, fmt.Sprintf("%d", lines[i]))
		} else {
			parts = append(parts, fmt.Sprintf("%d-%d", lines[i], lines[j]))
		}

		i = j + 1
	}

	return join.CommaSpace(parts)
}
