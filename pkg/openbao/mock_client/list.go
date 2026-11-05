package mock_client

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
	"strings"
)

func (c *Client) List(path string) ([]string, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	prefix := ""

	if path != "" {
		prefix = join.Empty(path, stringsConstant.Slash)
	}

	result := []string{}

	for stored := range c.secrets {
		rest, found := strings.CutPrefix(stored, prefix)

		if !found {
			continue
		}

		if head, _, folder := strings.Cut(rest, stringsConstant.Slash); folder {
			rest = join.Empty(head, stringsConstant.Slash)
		}

		if !slices.Contains(result, rest) {
			result = append(result, rest)
		}
	}

	slices.Sort(result)

	return result, nil
}
