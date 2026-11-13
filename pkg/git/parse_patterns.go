package git

import (
	"bufio"
	"bytes"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"strings"
)

func ParsePatterns(
	content []byte,
	domain []string,
) []gitignore.Pattern {
	var result []gitignore.Pattern
	s := bufio.NewScanner(bytes.NewReader(content))

	for s.Scan() {
		line := s.Text()

		if !strings.HasPrefix(line, stringsConstant.Hash) &&
			strings.TrimSpace(line) != "" {
			result = append(result, gitignore.ParsePattern(line, domain))
		}
	}

	return result
}
