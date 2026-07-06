package chunk

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"regexp"
	"strings"
)

var fencePattern = regexp.MustCompile(join.Empty(`\n`, "```"))

func findCodeFences(text string) []codeFence {
	matches := fencePattern.FindAllStringIndex(text, -1)
	var result []codeFence
	inFence := false
	start := 0

	for _, match := range matches {
		if !inFence {
			start = match[0]
			inFence = true
		} else {
			result = append(
				result,
				codeFence{
					start: start,
					end:   match[0] + len(strings.TrimSpace(text[match[0]:match[1]])),
				},
			)
			inFence = false
		}
	}

	if inFence {
		result = append(result, codeFence{start: start, end: len(text)})
	}

	return result
}
