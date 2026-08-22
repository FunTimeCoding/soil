package procfile

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"os"
	"strings"
)

func Parse(path string) ([]Entry, error) {
	content, e := os.ReadFile(path)

	if e != nil {
		return nil, e
	}

	var result []Entry

	for _, line := range strings.Split(string(content), "\n") {
		tokens := strings.SplitN(line, ":", 2)

		if len(tokens) != 2 || len(tokens[0]) == 0 || tokens[0][0] == '#' {
			continue
		}

		result = append(
			result,
			Entry{
				Name:    strings.TrimSpace(tokens[0]),
				Command: strings.TrimSpace(tokens[1]),
			},
		)
	}

	if len(result) == 0 {
		return nil, validation.New("no valid entries in %s", path)
	}

	return result, nil
}
