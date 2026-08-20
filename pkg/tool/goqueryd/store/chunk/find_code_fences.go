package chunk

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func findCodeFences(text string) []codeFence {
	matches := constant.FencePattern.FindAllStringIndex(text, -1)
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
					end: match[0] + len(
						strings.TrimSpace(text[match[0]:match[1]]),
					),
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
