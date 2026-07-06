package chunk

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk/break_pattern"
	"regexp"
	"sort"
)

var breakPatterns = []break_pattern.Pattern{
	{Pattern: regexp.MustCompile(`\n#{1}(?:[^#])`), Score: 100},
	{Pattern: regexp.MustCompile(`\n#{2}(?:[^#])`), Score: 90},
	{Pattern: regexp.MustCompile(`\n#{3}(?:[^#])`), Score: 80},
	{Pattern: regexp.MustCompile(`\n#{4}(?:[^#])`), Score: 70},
	{Pattern: regexp.MustCompile(`\n#{5}(?:[^#])`), Score: 60},
	{Pattern: regexp.MustCompile(`\n#{6}(?:[^#])`), Score: 50},
	{Pattern: regexp.MustCompile(join.Empty(`\n`, "```")), Score: 80},
	{Pattern: regexp.MustCompile(`\n(?:---|\*\*\*|___)\s*\n`), Score: 60},
	{Pattern: regexp.MustCompile(`\n\n+`), Score: 20},
	{Pattern: regexp.MustCompile(`\n[-*]\s`), Score: 5},
	{Pattern: regexp.MustCompile(`\n\d+\.\s`), Score: 5},
	{Pattern: regexp.MustCompile(`\n`), Score: 1},
}

func scanBreakPoints(text string) []breakPoint {
	seen := map[int]breakPoint{}

	for _, p := range breakPatterns {
		for _, match := range p.Pattern.FindAllStringIndex(text, -1) {
			position := match[0]
			existing, found := seen[position]

			if !found || p.Score > existing.score {
				seen[position] = breakPoint{
					position: position,
					score:    p.Score,
				}
			}
		}
	}

	result := make([]breakPoint, 0, len(seen))

	for _, b := range seen {
		result = append(result, b)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].position < result[j].position
		},
	)

	return result
}
