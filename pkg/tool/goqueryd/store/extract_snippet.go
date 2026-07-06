package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func ExtractSnippet(
	body string,
	query string,
	chunkPosition int,
) (string, int) {
	if body == "" {
		return "", 0
	}

	searchBody := body
	lineOffset := 0

	if chunkPosition > 0 {
		padding := 100
		start := chunkPosition - padding

		if start < 0 {
			start = 0
		}

		end := chunkPosition + constant.ChunkSize + padding

		if end > len(body) {
			end = len(body)
		}

		searchBody = body[start:end]

		if start > 0 {
			lineOffset = strings.Count(body[:start], separator.Unix)
		}
	}

	lines := strings.Split(searchBody, separator.Unix)
	terms := extractQueryTerms(query)
	bestLine := 0
	bestScore := -1

	for i, line := range lines {
		lower := strings.ToLower(line)
		score := 0

		for _, term := range terms {
			if strings.Contains(lower, term) {
				score++
			}
		}

		if score > bestScore {
			bestScore = score
			bestLine = i
		}
	}

	start := bestLine - 2

	if start < 0 {
		start = 0
	}

	end := bestLine + 3

	if end > len(lines) {
		end = len(lines)
	}

	snippet := strings.Join(lines[start:end], separator.Unix)

	if len(snippet) > constant.SnippetMaxLength {
		snippet = fmt.Sprintf(
			"%s...",
			snippet[:constant.SnippetMaxLength-3],
		)
	}

	return snippet, lineOffset + bestLine + 1
}
