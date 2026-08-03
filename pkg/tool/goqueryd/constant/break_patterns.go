package constant

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk/break_pattern"
	"regexp"
)

var BreakPatterns = []break_pattern.Pattern{
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
